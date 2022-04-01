package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rellyson/ifood-challenge/server/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Creating a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handlers.HealthcheckHandler(rr, req)
	res := rr.Result()

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Check the response body is what we expect.
	resBody, _ := io.ReadAll(res.Body)
	assert.JSONEq(t, `{"message": "OK"}`, string(resBody))
}
