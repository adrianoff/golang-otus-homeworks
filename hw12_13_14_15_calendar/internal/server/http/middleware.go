package internalhttp

import (
	"net/http"
)

func (s *Server) loggingMiddleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.RequestURI)
		s.ServeHTTP(w, r)
	})
}
