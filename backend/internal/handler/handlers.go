package handler

import (
	"net/http"
)

const (
	ctxKeyRequestID = iota
)

func (h *Handler) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.respondWithJSON(w, http.StatusOK, struct {
			Ping string `json:"ping"`
		}{
			Ping: "pong",
		})
	}
}
