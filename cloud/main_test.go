package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pc/cloud/Dto"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/cloud/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSetCluster(t *testing.T) {
	router := setupRouter()

	var clusterDto Dto.ClusterDto
	clusterDto.Name = "테스트"
	clusterDto.Type = "aws"
	clusterDto.Quntity = "1"
	pbytes, _ := json.Marshal(clusterDto)
	buff := bytes.NewBuffer(pbytes)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/v1/cloud/cluster", buff)
	req.Header.Add("Authorization", "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ3cmFHc0t4bUprb3M0Um44M3Ffc2RId251RnY0TllKbXliRjVMZllEOG5jIn0.eyJleHAiOjE1OTA3NDM3ODQsImlhdCI6MTU5MDc0MDE4NCwianRpIjoiMzg0ZDdiOWUtYTBlNS00OTI5LTkzYmMtMGU4NjQxOTRiYjhlIiwiaXNzIjoiaHR0cDovL2t1YmVybmV0ZXMuZG9ja2VyLmludGVybmFsOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwiYXVkIjpbIm1hc3Rlci1yZWFsbSIsImFjY291bnQiXSwic3ViIjoiZGJjZjQ5MjMtMzE4Zi00NzU2LWI1Y2QtZDU1ZTY3NmQzYjg0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZ29rdWJlciIsInNlc3Npb25fc3RhdGUiOiJlZGY3ZGJmZi0xODI2LTQ2MzYtOTQyNS1jYjQ3ZmY2MjdmY2UiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImNyZWF0ZS1yZWFsbSIsIm9mZmxpbmVfYWNjZXNzIiwiYWRtaW4iLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7Im1hc3Rlci1yZWFsbSI6eyJyb2xlcyI6WyJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctcmVhbG0iLCJtYW5hZ2UtaWRlbnRpdHktcHJvdmlkZXJzIiwiaW1wZXJzb25hdGlvbiIsImNyZWF0ZS1jbGllbnQiLCJtYW5hZ2UtdXNlcnMiLCJxdWVyeS1yZWFsbXMiLCJ2aWV3LWF1dGhvcml6YXRpb24iLCJxdWVyeS1jbGllbnRzIiwicXVlcnktdXNlcnMiLCJtYW5hZ2UtZXZlbnRzIiwibWFuYWdlLXJlYWxtIiwidmlldy1ldmVudHMiLCJ2aWV3LXVzZXJzIiwidmlldy1jbGllbnRzIiwibWFuYWdlLWF1dGhvcml6YXRpb24iLCJtYW5hZ2UtY2xpZW50cyIsInF1ZXJ5LWdyb3VwcyJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhZG1pbiJ9.UzbdhQi_KfGmq8CRr-ynW6q8rXi2cjt4bnWzD5SsnfdA6DiBDeQGgya9AdV5_yHBn903rBcbX6ZG7Rj_Pwvvq0rOMVAR45T3EJoUuNtDDQQbX1Nk8RqohONIlwM1lVI529s5wahurSjSuG-tahCmN27KOGYROjRelKIWSTr4K8ZtIZmKylR9dYNJervY96gMgOOPumlapc32X7xDYoPcgUKuRBylXhqR2_k9_rs5HWC8GMGjZVAK-watSuUNNzKHXButo2ryERZ8kd_ovvXPlF3saiVQvC5GtMEzOyChNqb-5pTh-34EuqJo6RRfnqimP-ro3JKCnKhVhTKQSjS7Lg")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	//resp, _ := http.Post("/v1/cloud/cluster", "application/json", buff)
	//
	//assert.Equal(t, 200, resp.Status)
}
