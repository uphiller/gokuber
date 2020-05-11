package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/vms", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 200, w.Body.String())
}

func TestPostResource(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"subscriptionId":"e8fe9bd4-424f-4247-b815-2c84e7c3d0a8",
							"clientId":"e350e1f4-9cda-4975-b8c3-4721dc24cb5b",
							"clientSecret":"sLEjVIVGLu]@olj/y9:xCq2OTta758cZ",
							"tenantId":"96fe0bf7-70bb-4a02-be1c-a37e898345ea"}`)
	req, _ := http.NewRequest("POST", "/v1/info", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 200, w.Body.String())
}
