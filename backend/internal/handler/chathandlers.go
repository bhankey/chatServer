package handler

import (
	"encoding/json"
	"net/http"
)

// @Summary new chat
// @Description creates new chat
// @Accept json
// @Produce json
// @Router /chats/add [post]
// @Param chat body models.AddChat true "Add chat"
// @Success 200 {object} models.ChatId chat_id
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

// @Summary get chat
// @Description get chat by user id
// @Accept json
// @Produce json
// @Router /chats/get [post]
// @Param chat body models.UserId true "user"
// @Success 200 {object} models.Chat
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
		h.respondWithJSON(w, http.StatusOK, chats)
	}
}
