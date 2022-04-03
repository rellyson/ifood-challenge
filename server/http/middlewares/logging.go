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
	RequestURI string              `json:"request_uri"`
	UserAgent  string              `json:"user_agent"`
	Headers    map[string][]string `json:"headers"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := LoggingPayload{}
		p.Date = time.Now()
		p.Method = r.Method
		p.RequestURI = r.RequestURI
		p.UserAgent = r.UserAgent()
		p.Headers = r.Header

		v, _ := json.Marshal(&p)
		fmt.Println(string(v))

		next.ServeHTTP(w, r)
	})
}
