package handler

import (
	"chatServer/internal/logger"
	"chatServer/internal/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	http.Handler
	logger  *logger.Logger
	service *service.Service
}

func NewHandler(s *service.Service, l *logger.Logger) (*Handler, error) {
	h := &Handler{
		logger:  l,
		service: s,
	}
	h.configureRoutes()
	return h, nil
}

func (h *Handler) configureRoutes() {
	r := mux.NewRouter()

	r.Use(h.setRequestID, h.logRequest)

	// installing CORS policy
	r.Use(
		handlers.CORS(
			handlers.AllowedHeaders([]string{"Content-Type"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodPost,
				http.MethodHead,
				http.MethodPut,
				http.MethodOptions,
			},
			),
		))
	h.configureUsersRoutes(r.PathPrefix("/users/").Subrouter())
	h.configureChatsRoutes(r.PathPrefix("/chats/").Subrouter())
	h.configureMessagesRoutes(r.PathPrefix("/messages/").Subrouter())
	h.Handler = r
}

func (h *Handler) configureUsersRoutes(r *mux.Router) {
	r.HandleFunc("/add", h.addUser()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/ping", h.ping())
}

func (h *Handler) configureChatsRoutes(r *mux.Router) {
	r.HandleFunc("/add", h.addChat()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/get", h.getChat()).Methods(http.MethodPost, http.MethodOptions)
}

func (h *Handler) configureMessagesRoutes(r *mux.Router) {
	r.HandleFunc("/add", h.addMessage()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/get", h.getMessages()).Methods(http.MethodPost, http.MethodOptions)
}
