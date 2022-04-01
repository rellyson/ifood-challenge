package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/rellyson/ifood-challenge/server/handlers"
)

func bootstrap() {
	log.Print("Bootstraping application...")
	a, exists := os.LookupEnv("APP_ADDR")

	if !exists {
		log.Fatal("Missing APP_ADDR environment variable!")
	}

	http.HandleFunc("/healthcheck", handlers.HealthcheckHandler)
	http.HandleFunc("/events/notify_alert", handlers.NotifyAlertHandler)

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
