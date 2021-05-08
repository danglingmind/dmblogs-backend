package middleware

import (
	"fmt"
	"net/http"

	"danglingmind.com/ddd/infrastructure/auth"
	"danglingmind.com/ddd/interfaces"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			interfaces.Error(w, http.StatusUnauthorized, err, err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if r.Method == "OPTIONS" {
			interfaces.Error(w, http.StatusNoContent, fmt.Errorf("CORS error"), "HTTP Method OPTIONS not allowed")
			return
		}
		next.ServeHTTP(w, r)
	})
}
