package server

import (
	"encoding/json"
	"net/http"
	"os"

	log "github.com/rs/zerolog/log"
)

type response struct {
	Hostname string `json:"hostname"`
}

func Server(address string) {
	HOSTNAME, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("path", "/").Str("method", r.Method).Msg("")
		json.NewEncoder(w).Encode(response{Hostname: HOSTNAME})
	})

	log.Info().Msgf("HTTP server listening on %s", address)
	http.ListenAndServe(address, nil)
}
