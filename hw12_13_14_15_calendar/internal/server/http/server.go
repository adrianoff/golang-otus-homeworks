package internalhttp

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/app"
	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	logger     logger.Logger
	app        app.Application
	address    string
	httpServer http.Server
}

func NewServer(logger logger.Logger, app app.Application, address string) *Server {
	return &Server{
		logger:  logger,
		app:     app,
		address: address,
	}
}

func (s *Server) Start(_ context.Context) error {
	handler := s.loggingMiddleware()
	s.httpServer = http.Server{
		Addr:         s.address,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	stopCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	s.logger.Info("server http is stopping...")
	if err := s.httpServer.Shutdown(stopCtx); err != nil {
		s.logger.Error("failed to stop http server: " + err.Error())
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode("hello world"); err != nil {
			http.Error(w, "Encode error", 500)
		}
	case "/create/":
		// Just for testing
		if err := s.create(w, r); err != nil {
			http.Error(w, "Create error", 500)
		}
	default:
		http.NotFound(w, r)
	}
}

func (s *Server) create(w http.ResponseWriter, _ *http.Request) error { //nolint:unparam
	// Just for testing
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()
	id, _ := s.app.CreateEvent(ctx, "TEST")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, "Encode error", 500)
	}

	return nil
}
