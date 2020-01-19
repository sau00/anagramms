package anagram

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"gitlab.com/anagramms/handler"
	"gitlab.com/anagramms/storage"
)

const (
	version = "0.1"
)

type Handler struct {
	handler.DefaultHandler
}

func (h *Handler) Version(ctx *fasthttp.RequestCtx) {
	h.JSON(ctx, http.StatusOK, VersionResponse{
		Version: version,
		Message: "Hello & Welcome to fastanagram API!",
	})
}

type Database struct {
	Anagrams map[string][]string
}

func (h *Handler) Get(ctx *fasthttp.RequestCtx) {
	req := GetRequest{
		Word: string(ctx.FormValue("word")),
	}

	s := storage.Get()
	s.MTX.RLock()
	defer s.MTX.RUnlock()

	anagram := storage.TAnagram(req.Word)
	anagrams := s.AnagramsR[anagram.Key()]

	h.JSON(ctx, http.StatusOK, GetResponse{
		Len:      len(anagrams),
		Anagrams: anagrams,
	})
}

func (h *Handler) Load(ctx *fasthttp.RequestCtx) {
	req := LoadRequest{}

	if err := json.Unmarshal(ctx.PostBody(), &req.Anagrams); err != nil {
		h.Error(ctx, http.StatusInternalServerError, err)
	}

	s := storage.Get()
	s.MTX.RLock()
	defer s.MTX.RUnlock()

	for _, anagram := range req.Anagrams {
		// anagram := storage.TAnagram(a)
		key := anagram.Key()

		// Если слов с такой комбинацией букв не встречалось, создаем запись в хранилище
		if _, ok := s.AnagramsW[key]; !ok {
			s.AnagramsW[key] = make(map[storage.TAnagram]bool, len(req.Anagrams))
			s.AnagramsR[key] = make([]storage.TAnagram, 0, len(req.Anagrams))
		}

		// Проверяем, встречалась ли данная конкретная анаграмма, если нет сохраняем
		if _, ok := s.AnagramsW[key][anagram]; !ok {
			s.AnagramsW[key][anagram] = true
			s.AnagramsR[key] = append(s.AnagramsR[key], anagram)
		}
	}

	log.Info().Int("Len", len(req.Anagrams)).Msg("anagrams received")

	h.JSON(ctx, http.StatusOK, LoadResponse{
		Len:     len(req.Anagrams),
		Message: "success",
	})
}
