package ginsrv

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type response struct {
	Hostname string `json:"hostname"`
}

func Server(address string) {
	HOSTNAME, _ := os.Hostname()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response{Hostname: HOSTNAME})
	})

	fmt.Println("Gin listening on " + address)
	r.Run(address)
}
