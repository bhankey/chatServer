package handler

import (
	"chatServer/internal/logger"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// w http.ResponseWriter, r *http.Request, statusCode int, err error
func Test_respondWithError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		data       error
		json       []byte
	}{
		{
			name:       "with some client error",
			statusCode: 404,
			data:       fmt.Errorf("some client error"),
			json:       []byte("{\"error\":\"some client error\"}\n"),
		},
		{
			name:       "with some server error",
			statusCode: 500,
			data:       fmt.Errorf("some server error"),
			json:       []byte("{\"error\":\"internal server error\"}\n"),
		},
	}
	l, _ := logger.NewLogger(logger.NewConfig())
	h, err := NewHandler(nil, l)
	assert.NoError(t, err)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRecorder()
			w := httptest.NewRequest("POST", "/ping", nil)
			h.respondWithError(r, w, test.statusCode, test.data)
			assert.Equal(t, r.Code, test.statusCode)
			assert.Equal(t, r.Body.Bytes(), test.json)
		})
	}

}

func Test_respondWithJSON(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		data       interface{}
		json       []byte
	}{
		{
			name:       "empty body",
			statusCode: 200,
			data:       nil,
			json:       nil,
		},
		{
			name:       "404 and body with body",
			statusCode: 404,
			data:       map[string]string{"error": "someError"},
			json:       []byte("{\"error\":\"someError\"}\n"),
		},
		{
			name:       "200 and body with some data",
			statusCode: 200,
			data:       map[string]string{"abc": "abc", "somedata": "somedata"},
			json:       []byte("{\"abc\":\"abc\",\"somedata\":\"somedata\"}\n"),
		},
	}
	h, err := NewHandler(nil, nil)
	assert.NoError(t, err)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRecorder()
			h.respondWithJSON(r, test.statusCode, test.data)
			assert.Equal(t, r.Code, test.statusCode)
			assert.Equal(t, r.Body.Bytes(), test.json)
		})
	}
}
