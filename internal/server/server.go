package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/YaroslavGaponov/whereami/internal/whereami"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	port     int
	router   *chi.Mux
	whereAmI *whereami.WhereAmI
}

func New(port int, whereAmI *whereami.WhereAmI) Server {

	server := Server{
		port:     port,
		router:   chi.NewRouter(),
		whereAmI: whereAmI,
	}

	server.router.Get("/whereami", server.SearchHandler)

	return server
}

func (server *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", server.port), server.router)
}

func (server *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {

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
