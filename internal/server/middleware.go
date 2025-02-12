package server

import (
	"net/http"
	"strings"

	"github.com/theo-krutiy/minimal-go/internal/auth"
)

func (s *Server) requiresLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if err := auth.ValidateToken(token, s.secret); err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
