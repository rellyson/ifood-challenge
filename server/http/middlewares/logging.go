package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LoggingPayload struct {
	Date       time.Time           `json:"date"`
	Method     string              `json:"method"`
	RequestURL string              `json:"request_url"`
	UserAgent  string              `json:"user_agent"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := LoggingPayload{}
		p.Date = time.Now()
		p.Method = r.Method
		p.RequestURL = r.RequestURI
		p.UserAgent = r.UserAgent()
		p.Headers = r.Header
		p.Body = convertBodyToString(r.Body)

		v, _ := json.Marshal(&p)
		fmt.Println(string(v))
	})
}

func convertBodyToString(b io.ReadCloser) string {
	body, _ := io.ReadAll(b)

	return fmt.Sprintf("%v", string(body))
}
