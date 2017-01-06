package filter

import (
	"net/http"
	logger "github.com/Sirupsen/logrus"
	"github.com/fatih/stopwatch"
)

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := stopwatch.Start(0)
		h.ServeHTTP(w, r)
		logger.Info("[", r.Method, "] - ", r.URL, " Nanoseconds=", s.ElapsedTime().Nanoseconds())
	})
}