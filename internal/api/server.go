package api

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/DHSY-ishere/kanbanAPI/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	cfg    config.Config
	db     *sql.DB
	router *chi.Mux
}

func NewServer(cfg config.Config, db *sql.DB) *Server {
	s := &Server{
		cfg:    cfg,
		db:     db,
		router: chi.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", s.handleHealth)
		r.Get("/version", s.handleVersion)

	})
}

func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"version":"0.0.1"}`))
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := s.db.PingContext(ctx); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status":"db down"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.cfg.Addr, s.router)
}
