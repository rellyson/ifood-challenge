package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rellyson/ifood-challenge/server/handlers"
)

func bootstrap() {
	log.Print("Bootstraping application...")
	a, exists := os.LookupEnv("APP_ADDR")

	if !exists {
		log.Fatal("Missing APP_ADDR environment variable!")
	}

	router := mux.NewRouter()
	s := router.PathPrefix("/v1/").Subrouter()

	s.HandleFunc("/healthcheck", handlers.HealthcheckHandler).Methods(http.MethodGet)
	s.HandleFunc("/events/notify_alert", handlers.NotifyAlertHandler).Methods(http.MethodPost).Headers("Content-type", "application/json")

	http.Handle("/", s)
	listener, err := net.Listen("tcp", a)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Application is listenning at: %v", a)
	log.Fatal(http.Serve(listener, nil))
}

func main() {
	bootstrap()
}
