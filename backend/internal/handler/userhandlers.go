package handler

import (
	"encoding/json"
	"net/http"
)

// @Summary new user
// @Description creates new user
// @Accept json
// @Produce json
// @Router /users/add [post]
// @Param chat body models.AddUser true "Add user"
// @Success 200 {object} models.UserId user_id
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
