package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/theo-krutiy/minimal-go/internal/auth"
	"github.com/theo-krutiy/minimal-go/internal/models"
	"github.com/theo-krutiy/minimal-go/internal/shop"
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

func (s *Server) handleGetItems() http.HandlerFunc {
	type response struct {
		Page       []models.Item `json:"page"`
		TotalCount int           `json:"totalCount"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(query.Get("offset"))
		if err != nil {
			offset = 0
		}
		limit, err := strconv.Atoi((query.Get("limit")))
		if err != nil {
			limit = 10
		}
		q := query.Get("query")

		items, totalCount, err := shop.GetItems(q, offset, limit, s.Db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		page := make([]models.Item, len(items))
		for i, item := range items {
			page[i] = *item
		}

		res := response{page, totalCount}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}

	}
}
