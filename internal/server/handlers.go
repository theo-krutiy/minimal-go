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
