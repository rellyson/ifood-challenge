package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rellyson/ifood-challenge/server/utils"
	"github.com/stretchr/testify/assert"
)

func TestServerError(t *testing.T) {
	rr := httptest.NewRecorder()

	errorMessage := "Testing error"
	e := utils.ServerError(rr, http.StatusBadRequest, errorMessage)

	assert.Equal(t, http.StatusBadRequest, e.Status)
	assert.Equal(t, errorMessage, e.Error)
}
