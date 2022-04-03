package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rellyson/ifood-challenge/server/aws"
	"github.com/rellyson/ifood-challenge/server/controllers"
	"github.com/rellyson/ifood-challenge/server/http/middlewares"
	"github.com/rellyson/ifood-challenge/server/service"
)

var (
	sqsClient             aws.SQSClient                     = aws.NewSQSClient()
	messaging             service.MessagingService          = service.NewMessagingService(sqsClient)
	healthcheckController controllers.HealthCheckController = controllers.NewHealthCheckController()
	eventsController      controllers.EventsController      = controllers.NewEventsController(messaging)
)

func UseRouter(router *mux.Router) {
	s := router.PathPrefix("/v1/").Subrouter()
	s.Use(middlewares.LoggingMiddleware)

	//setting routes
	s.HandleFunc("/healthcheck", healthcheckController.Status).Methods(http.MethodGet)
	s.HandleFunc("/events/notify_alert", eventsController.NotifyAlert).Methods(http.MethodPost).Headers("Content-type", "application/json")
}
