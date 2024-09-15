package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YaroslavGaponov/whereami/internal/whereami"
	"github.com/YaroslavGaponov/whereami/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	ctx      context.Context
	addr     string
	router   *chi.Mux
	whereAmI *whereami.WhereAmI
}

func New(ctx context.Context, addr string, whereAmI *whereami.WhereAmI) Server {
	server := Server{
		ctx:      ctx,
		addr:     addr,
		router:   chi.NewRouter(),
		whereAmI: whereAmI,
	}

	server.router.Get("/whereami", server.searchHandler)
	server.router.Get("/alive", server.alive)
	server.router.Get("/ready", server.ready)

	return server
}

func (server *Server) Run() error {
	logger.GetLogger(server.ctx).Info("server is starting...")
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

	point, err := server.whereAmI.Search(flat, flng)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(point); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func (server *Server) alive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (server *Server) ready(w http.ResponseWriter, r *http.Request) {
	initialized := server.whereAmI.IsInitialized()
	if initialized {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
