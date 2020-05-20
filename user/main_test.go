package main

import (
	"encoding/json"
	"testing"

	// REST
	"gopkg.in/resty.v1"

	// Encoding
	b64 "encoding/base64"
)

func getBasicAuthForClient(clientId string, clientSecret string) string {
	var httpBasicAuth string
	if len(clientId) > 0 && len(clientSecret) > 0 {
		httpBasicAuth = b64.URLEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	}

	return "Basic " + httpBasicAuth
}

func TestPingRoute(t *testing.T) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", getBasicAuthForClient("gokubermanagement-user", "b53f3ca6-0559-4781-a9ee-a35cb0d7a64a")).
		SetFormData(map[string]string{
			"grant_type": "password",
			"username":   "tester",
			"password":   "1234",
		}).Post("http://localhost:8080/auth/realms/master/protocol/openid-connect/token")
	if err != nil {
		print(err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {

	}
	print(result["access_token"])
}
