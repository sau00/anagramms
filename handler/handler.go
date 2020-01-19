package handler

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

const (
	_headerContentType     = "Content-Type"
	_headerApplicationJSON = "application/json"
)

type DefaultHandler struct {
}

func (h *DefaultHandler) JSON(ctx *fasthttp.RequestCtx, code int, resp interface{}) {
	ctx.Response.Header.SetCanonical(
		[]byte(_headerContentType),
		[]byte(_headerApplicationJSON),
	)

	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(resp); err != nil {
		log.Error().Err(err).Msg("")
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

type errorResponse struct {
	Code    int
	Message string
}

func (h *DefaultHandler) Error(ctx *fasthttp.RequestCtx, code int, err error) {
	ctx.Response.Header.SetCanonical(
		[]byte(_headerContentType),
		[]byte(_headerApplicationJSON),
	)

	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(errorResponse{
		Code:    code,
		Message: err.Error(),
	}); err != nil {
		log.Error().Err(err).Msg("")
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
