package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) addChat() http.HandlerFunc {
	type request struct {
		Name    string `json:"name"`
		UsersId []int  `json:"users"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.respondWithError(w, r, http.StatusBadRequest, err)
			return
		}
		chatId, err := h.service.Chat.Create(req.Name, req.UsersId)
		if err != nil {
			h.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"chat_id": chatId,
		})
	}
}

func (h *Handler) getChat() http.HandlerFunc {
	type request struct {
		UserId int `json:"user"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.respondWithError(w, r, http.StatusBadRequest, err)
			return
		}
		chats, err := h.service.Chat.GetByUserId(req.UserId)
		if err != nil {
			h.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"chats": chats,
		})
	}
}
