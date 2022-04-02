package main

import (
	"log"

	"github.com/rellyson/ifood-challenge/server/aws"
	"github.com/rellyson/ifood-challenge/server/http"
)

func bootstrap() {
	log.Print("Bootstraping application...")

	//set and start a new SQSConfig
	aws.NewSQSClient()

	// set configs and start http server
	http.SetHttpServer()
}

func main() {
	bootstrap()
}
