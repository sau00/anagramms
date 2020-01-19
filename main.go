package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"gitlab.com/anagramms/handler/anagram"
	"gitlab.com/anagramms/storage"
)

func main() {
	storage.Get()

	router := fasthttprouter.New()
	{
		h := anagram.Handler{}
		router.GET("/", h.Version)
		router.POST("/load", h.Load)
		router.GET("/get", h.Get)
	}

	log.Info().Msg("starting web service")
	if err := fasthttp.ListenAndServe(":8080", router.Handler); err != nil {
		log.Fatal().Err(err).Msg("failed to start web service")
	}
}
