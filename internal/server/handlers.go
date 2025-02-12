package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/theo-krutiy/minimal-go/internal/auth"
)

func (s *Server) handleCreateUser() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	type response struct {
		UserId string `json:"UserId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Couldn't parse request", http.StatusBadRequest)
			return
		}
		newUserId, err := auth.CreateNewUser(req.Login, req.Password, s.Db)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err.Error()), http.StatusInternalServerError)
			return
		}

		res := response{
			UserId: newUserId,
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Couldn't unmarshall response", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleAuthenticate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	type response struct {
		Token string `json:"token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("%v", err.Error()), http.StatusInternalServerError)
			return
		}

		token, err := auth.Authenticate(req.Login, req.Password, s.secret, s.Db)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		res := response{Token: token}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Couldn't unmarshall response", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleProtectedRoute(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "this is a protected resource")
}
