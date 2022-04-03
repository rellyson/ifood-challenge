package http

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func SetHttpServer() {
	a, exists := os.LookupEnv("APP_ADDR")

	if !exists {
		log.Fatal("Missing APP_ADDR environment variable!")
	}

	router := mux.NewRouter()
	UseRouter(router)

	log.Printf("Starting to serve application at: %v", a)
	log.Fatal(http.ListenAndServe(a, router))
}
