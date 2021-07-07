package handler

import (
	"encoding/json"
	"net/http"
)

// @Summary new message
// @Description sends new message to chat from user
// @Accept json
// @Produce json
// @Router /messages/add [post]
// @Param chat body models.AddMessage true "Add message"
// @Success 200 {object} models.MessageId message_id
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

// @Summary get messages
// @Description get all messages from chat
// @Accept json
// @Produce json
// @Router /messages/get [post]
// @Param chat body models.ChatIdS true "chat id"
// @Success 200 {object} []models.Message messages
func (h *Handler) getMessages() http.HandlerFunc {
	type request struct {
		ChatId int `json:"chat"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.respondWithError(w, r, http.StatusBadRequest, err)
			return
		}
		messages, err := h.service.Message.GetByChatId(req.ChatId)
		if err != nil {
			h.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respondWithJSON(w, http.StatusOK, messages)
	}
}
