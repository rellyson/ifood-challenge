package controllers_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rellyson/ifood-challenge/server/controllers"
	"github.com/rellyson/ifood-challenge/server/controllers/mocks"
	"github.com/rellyson/ifood-challenge/server/service"
	"github.com/stretchr/testify/assert"
)

func TestNotifyAlert(t *testing.T) {
	serviceMock := new(mocks.MessageService)
	testController := controllers.NewEventsController(serviceMock)

	serviceMock.On("SendMessageToQueue").Return(service.SendMessageResponse{}, nil)

	req, err := http.NewRequest("POST", "/v1/events/notify-alert", strings.NewReader(`{"channel": "test", "message": "test"}`))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testController.NotifyAlert(rr, req)
	res := rr.Result()

	expectedPayload := `{
		"md5_of_body":"",
		"message_id":"",
		"status":""
	}`

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	// Check body payload
	body, _ := io.ReadAll(res.Body)
	assert.JSONEq(t, expectedPayload, string(body))
}

func TestPayloadValidation(t *testing.T) {
	serviceMock := new(mocks.MessageService)
	testController := controllers.NewEventsController(serviceMock)

	req, err := http.NewRequest("POST", "/v1/events/notify-alert", strings.NewReader(`{"channel": "test"}`))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testController.NotifyAlert(rr, req)
	res := rr.Result()

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	// Check body payload
	expectedPayload := `{
		"status": 400,
		"error": "Payload is invalid: message field is missing"
	}`

	body, _ := io.ReadAll(res.Body)
	assert.JSONEq(t, expectedPayload, string(body))
}

func TestServiceErrorHandling(t *testing.T) {
	serviceMock := new(mocks.MessageService)
	testController := controllers.NewEventsController(serviceMock)

	serviceMock.On("SendMessageToQueue").Return(service.SendMessageResponse{}, errors.New("test error"))

	req, err := http.NewRequest("POST", "/v1/events/notify-alert", strings.NewReader(`{"channel": "test", "message": "test"}`))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testController.NotifyAlert(rr, req)
	res := rr.Result()

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	// Check the status code is what we expect.
	expectedPayload := `{
		"status": 500,
		"error": "Something went wrong. try again later"
	}`

	body, _ := io.ReadAll(res.Body)
	assert.JSONEq(t, expectedPayload, string(body))
}
