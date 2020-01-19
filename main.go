package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"gitlab.com/anagramms/handler/anagram"
	"gitlab.com/anagramms/storage"
)

func main() {
	router := fasthttprouter.New()

	storage.Get()

	{
		h := anagram.Handler{}
		router.GET("/", h.Version)
		router.POST("/load", h.Load)
		router.GET("/get", h.Get)
	}

	// router.GET("/get", Get)
	// router.POST("/load", Load)

	log.Info().Msg("starting web service")
	if err := fasthttp.ListenAndServe(":8080", router.Handler); err != nil {
		log.Fatal().Err(err).Msg("failed to start service")
	}
}
