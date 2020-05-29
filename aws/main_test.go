package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/aws/cluster", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
