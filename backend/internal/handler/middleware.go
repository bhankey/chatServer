package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// setRequestID set UUID for every request for logging
func (h *Handler) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewRandom()
		if err != nil {
			h.logger.Warnf("uuid creation failed: %v", err)
			return
		}
		sId := id.String()
		w.Header().Set("Request-ID", sId) // Not recommended to start custom headers with X- anymore
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, sId)))
	})
}

// logRequest log all requests
func (h *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := h.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		rw := NewResponseWriter(w)
		start := time.Now()

		next.ServeHTTP(rw, r)

		logger.Infof("completed with %d %s in %v", rw.statusCode, http.StatusText(rw.statusCode), time.Now().Sub(start))
	})
}
