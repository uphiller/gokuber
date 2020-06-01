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

var authorization = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ3cmFHc0t4bUprb3M0Um44M3Ffc2RId251RnY0TllKbXliRjVMZllEOG5jIn0.eyJleHAiOjE1OTA5Nzg4NDMsImlhdCI6MTU5MDk3NTI0MywianRpIjoiNDlhZmY5ZDctMTAwYi00MjRmLTg2YjctODcwZDNlNTI4ZjE0IiwiaXNzIjoiaHR0cDovL2t1YmVybmV0ZXMuZG9ja2VyLmludGVybmFsOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwiYXVkIjpbIm1hc3Rlci1yZWFsbSIsImFjY291bnQiXSwic3ViIjoiZGJjZjQ5MjMtMzE4Zi00NzU2LWI1Y2QtZDU1ZTY3NmQzYjg0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZ29rdWJlciIsInNlc3Npb25fc3RhdGUiOiI2NmYzNDk5Zi1mZjcwLTQ0MWMtYTJiMi1iMDc3MzJkYTZmYWIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImNyZWF0ZS1yZWFsbSIsIm9mZmxpbmVfYWNjZXNzIiwiYWRtaW4iLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7Im1hc3Rlci1yZWFsbSI6eyJyb2xlcyI6WyJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctcmVhbG0iLCJtYW5hZ2UtaWRlbnRpdHktcHJvdmlkZXJzIiwiaW1wZXJzb25hdGlvbiIsImNyZWF0ZS1jbGllbnQiLCJtYW5hZ2UtdXNlcnMiLCJxdWVyeS1yZWFsbXMiLCJ2aWV3LWF1dGhvcml6YXRpb24iLCJxdWVyeS1jbGllbnRzIiwicXVlcnktdXNlcnMiLCJtYW5hZ2UtZXZlbnRzIiwibWFuYWdlLXJlYWxtIiwidmlldy1ldmVudHMiLCJ2aWV3LXVzZXJzIiwidmlldy1jbGllbnRzIiwibWFuYWdlLWF1dGhvcml6YXRpb24iLCJtYW5hZ2UtY2xpZW50cyIsInF1ZXJ5LWdyb3VwcyJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhZG1pbiJ9.flvCg36gLN1Aa5Njxi2qOo7ufJPC46P_TUp0Rx0qWuGFedYioHFh-xUpLqyKEFPR2BB3ENjMNCpaWtVJZM79fOxKLtPgbTdA9gD9dhl90MwPfVA7aOwi-enDOBT2FP2okPhaGJRPRjNVzKIZhib-exFnpGLAEUz4xopAs5ANRtC0VgXw8v-83GYL6III3lThRbTNPDHrjNaCKleZotBZ6dA1Z534Plp_D-CPAyhpx7kD7PUAD3cBL36OxPeEyWTbCLdCniKXjZKMIMWoA_Rf4QorXM6TGf_vB2YQr_RCARVcwCNT-P5sr9TWbbWq7TH-8EsI56-ikD64lwvvGDGfVg"

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
	secretDto.Access_id = "123"
	secretDto.Secret_key = "123"
	pbytes, _ := json.Marshal(secretDto)
	buff := bytes.NewBuffer(pbytes)

	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/v1/cloud/secret", buff)
	req.Header.Add("Authorization", authorization)
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)

	//resp, _ := http.Post("/v1/cloud/cluster", "application/json", buff)
	//
	//assert.Equal(t, 200, resp.Status)
}
