package middlewares

import (
	"net/http"
)

func HeaderInjection(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Cache-Control", "public max-age=86400")
		w.Header().Set("Vary: origin", "origin")
		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}
