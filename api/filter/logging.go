package filter

import (
	"net/http"
	logger "github.com/Sirupsen/logrus"
)

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("[",r.Method,"] - ",r.URL,"")
		h.ServeHTTP(w, r)
	})
}