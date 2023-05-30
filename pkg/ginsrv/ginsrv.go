package ginsrv

import (
	"net/http"
	"os"

	log "github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type response struct {
	Hostname string `json:"hostname"`
}

func Server(address string) {
	HOSTNAME, _ := os.Hostname()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(ctx *gin.Context) {
		log.Info().Str("path", ctx.Request.RequestURI).Str("method", ctx.Request.Method).Msg("")
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response{Hostname: HOSTNAME})
	})

	log.Info().Msgf("Gin listening on %s", address)
	r.Run(address)
}
