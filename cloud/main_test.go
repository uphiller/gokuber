package main

import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pc/cloud/Config"
	"pc/cloud/Dto"
	"testing"
)

var authorization = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ3cmFHc0t4bUprb3M0Um44M3Ffc2RId251RnY0TllKbXliRjVMZllEOG5jIn0.eyJleHAiOjE1OTA5OTMzNDcsImlhdCI6MTU5MDk4OTc0NywianRpIjoiZjQzOTY2ZTYtNzAzZS00OWQyLTk5NjItODA2ZTRkOTYzOWRjIiwiaXNzIjoiaHR0cDovL2t1YmVybmV0ZXMuZG9ja2VyLmludGVybmFsOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwiYXVkIjpbIm1hc3Rlci1yZWFsbSIsImFjY291bnQiXSwic3ViIjoiZGJjZjQ5MjMtMzE4Zi00NzU2LWI1Y2QtZDU1ZTY3NmQzYjg0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZ29rdWJlciIsInNlc3Npb25fc3RhdGUiOiIyM2NjMTZhOC0wODAwLTQ4YTAtOGMxZC1lYTFmM2IwOTY2NzMiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImNyZWF0ZS1yZWFsbSIsIm9mZmxpbmVfYWNjZXNzIiwiYWRtaW4iLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7Im1hc3Rlci1yZWFsbSI6eyJyb2xlcyI6WyJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctcmVhbG0iLCJtYW5hZ2UtaWRlbnRpdHktcHJvdmlkZXJzIiwiaW1wZXJzb25hdGlvbiIsImNyZWF0ZS1jbGllbnQiLCJtYW5hZ2UtdXNlcnMiLCJxdWVyeS1yZWFsbXMiLCJ2aWV3LWF1dGhvcml6YXRpb24iLCJxdWVyeS1jbGllbnRzIiwicXVlcnktdXNlcnMiLCJtYW5hZ2UtZXZlbnRzIiwibWFuYWdlLXJlYWxtIiwidmlldy1ldmVudHMiLCJ2aWV3LXVzZXJzIiwidmlldy1jbGllbnRzIiwibWFuYWdlLWF1dGhvcml6YXRpb24iLCJtYW5hZ2UtY2xpZW50cyIsInF1ZXJ5LWdyb3VwcyJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhZG1pbiJ9.NoK2COjgREqTgrzfALh3ArF5IKvNbAjKM5WSaw7ePsR6YoCuelpPANryEJAwAHY2pz80ULayj0jxWAFcTRupwCDw_f09sURSj_NXtgVqDWBAyQ5G1OsUguR4Dsl6lBGjtISD9c3m-JQRxADiLQ7SUTUUFmdKsE6cL42pl527StnQB2hxIqD2u_MRcDsCWmJ_Tz1PEP0XIhy9bccAUZJ4l9iQPygbYugtYs_bj0z8tcN-VJ1FU7LsC8gBCfAIrVBzDuZMHHbbDAretGZLhLQlBro6Gi8CNR8HCMxtiiwaeMiTh2UB156KDOzRsiX-W8cmOae-ounyw0duG5H4W6lnUA"

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/cloud/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func initDatabase() {
	Config.DB, _ = gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/gokuber?charset=utf8&parseTime=True&loc=Local")
	Config.DB.LogMode(true)
}

func TestSetCluster(t *testing.T) {
	initDatabase()
	router := setupRouter()

	var clusterDto Dto.ClusterDto
	clusterDto.Name = "테스트"
	clusterDto.Type = "aws"
	clusterDto.Quntity = "1"
	clusterDto.Secret_id = 1
	pbytes, _ := json.Marshal(clusterDto)
	buff := bytes.NewBuffer(pbytes)

	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/v1/cloud/cluster", buff)
	req.Header.Add("Authorization", authorization)
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestSetSecret(t *testing.T) {
	initDatabase()
	router := setupRouter()

	var secretDto Dto.SecretDto
	secretDto.Name = "테스트"
	secretDto.Type = "aws"
	secretDto.Access_id = "AKIAS2KAI3ZLA5TJPI6N"
	secretDto.Secret_key = "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"
	pbytes, _ := json.Marshal(secretDto)
	buff := bytes.NewBuffer(pbytes)

	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/v1/cloud/secret", buff)
	req.Header.Add("Authorization", authorization)
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}
func TestGetSecret(t *testing.T) {
	initDatabase()
	router := setupRouter()
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/cloud/secrets?type=aws", nil)
	req.Header.Add("Authorization", authorization)
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}
