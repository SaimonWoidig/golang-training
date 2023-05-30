package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type response struct {
	Hostname string `json:"hostname"`
}

func Server(address string) {
	HOSTNAME, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(response{Hostname: HOSTNAME})
	})

	fmt.Println("Listening on " + address)
	http.ListenAndServe(address, nil)
}
