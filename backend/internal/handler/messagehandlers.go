package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) addMessage() http.HandlerFunc {
	type request struct {
		ChatId int    `json:"chat"`
		UserId int    `json:"author"`
		Text   string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.respondWithError(w, r, http.StatusBadRequest, err)
			return
		}
		messageId, err := h.service.Message.Create(req.ChatId, req.UserId, req.Text)
		if err != nil {
			h.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"message_id": messageId,
		})
	}
}
