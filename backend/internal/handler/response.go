package handler

import (
	"chatServer/pkg/db/sqlstore/postgres"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// respondWithError call if error happened in request processing
func (h *Handler) respondWithError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {

	if pgErr := postgres.PgError(err); pgErr != nil {
		err = pgErr
		statusCode = http.StatusForbidden
	}

	if statusCode < 500 {
		h.respondWithJSON(w, statusCode, map[string]string{"error": err.Error()})
	} else {
		logger := h.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Warnf("Server error: %v", err)
		h.respondWithJSON(w, statusCode, map[string]string{"error": "internal server error"})
	}
}

// respondWithJSON writes status code and json in response
func (h *Handler) respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
