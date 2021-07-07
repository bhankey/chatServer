package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) addUser() http.HandlerFunc {
	type request struct {
		UserName string `json:"username"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.respondWithError(w, r, http.StatusBadRequest, err)
			return
		}
		userId, err := h.service.User.Create(req.UserName)
		if err != nil {
			h.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"user_id": userId,
		})
	}
}
