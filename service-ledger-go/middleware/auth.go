package middleware

import (
	"net/http"
	"strings"

	"ledger-service/auth"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "missing auth header", http.StatusUnauthorized)
			return
		}
	
		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := auth.VerifyToken(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-User-ID", claims["userId"].(string))
		next.ServeHTTP(w, r)
	})
}