package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rellyson/ifood-challenge/server/handlers"
	"github.com/stretchr/testify/assert"
)

func TestNotifyAlertHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/events/notify-alert", strings.NewReader(`{"channel": "test", "message": "test"}`))

	if err != nil {
		t.Fatal(err)
	}

	// Creating a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handlers.NotifyAlertHandler(rr, req)
	res := rr.Result()

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestNotifyAlertHandlerPayloadValidation(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/events/notify-alert", strings.NewReader(`{"channel": "test"}`))

	if err != nil {
		t.Fatal(err)
	}

	// Creating a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handlers.NotifyAlertHandler(rr, req)
	res := rr.Result()

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
