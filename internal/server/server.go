package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YaroslavGaponov/whereami/internal/whereami"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	addr     string
	router   *chi.Mux
	whereAmI *whereami.WhereAmI
}

func New(addr string, whereAmI *whereami.WhereAmI) Server {

	server := Server{
		addr:     addr,
		router:   chi.NewRouter(),
		whereAmI: whereAmI,
	}

	server.router.Get("/whereami", server.searchHandler)

	return server
}

func (server *Server) Run() error {
	return http.ListenAndServe(server.addr, server.router)
}

func (server *Server) searchHandler(w http.ResponseWriter, r *http.Request) {

	lat := r.URL.Query().Get("lat")
	flat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	lng := r.URL.Query().Get("lng")
	flng, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	point := server.whereAmI.Search(flat, flng)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(point); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
