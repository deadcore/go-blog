package filter

import (
	"net/http"
)

func JsonContentTypeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Header().Set("Content-Type", "application/json")
	})
}