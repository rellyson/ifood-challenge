package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LoggingPayload struct {
	Date       time.Time           `json:"date"`
	Method     string              `json:"method"`
	RequestURL string              `json:"request_url"`
	UserAgent  string              `json:"user_agent"`
	Headers    map[string][]string `json:"headers"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := LoggingPayload{}
		p.Date = time.Now()
		p.RequestURL = r.RequestURI
		p.Headers = r.Header
		p.Method = r.Method
		p.UserAgent = r.UserAgent()

		v, _ := json.Marshal(&p)
		fmt.Println(string(v))
	})
}
