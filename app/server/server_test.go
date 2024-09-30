package server

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/app/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer(t *testing.T) {
	s := NewAPIServer(config.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello World!")
}
