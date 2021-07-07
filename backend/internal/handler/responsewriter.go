package handler

import "net/http"

// Idea from https://ndersson.me/post/capturing_status_code_in_net_http/

// ResponseWriter struct that helps to log status code of response
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewResponseWriter new ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		w,
		http.StatusOK,
	}
}

// WriteHeader expands original function for saving status code of response
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
