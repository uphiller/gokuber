package main

import (
	"encoding/json"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"testing"

	// REST
	"gopkg.in/resty.v1"

	// Encoding
	b64 "encoding/base64"
	"golang.org/x/net/context"
	_ "golang.org/x/oauth2"
)

func getBasicAuthForClient(clientId string, clientSecret string) string {
	var httpBasicAuth string
	if len(clientId) > 0 && len(clientSecret) > 0 {
		httpBasicAuth = b64.URLEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	}

	return "Basic " + httpBasicAuth
}

func TestLoginPassword(t *testing.T) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "password",
			"username":      "admin",
			"password":      "admin",
			"client_id":     "gokuber-user",
			"client_secret": "7a0cc26b-1bae-43b1-8c7f-15676108faab",
		}).Post("http://kubernetes.docker.internal/auth/realms/master/protocol/openid-connect/token")
	if err != nil {
		print(err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {

	}
	print(result["access_token"])
}

func TestLoginClientCredential(t *testing.T) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/master")
	if err != nil {
		log.Fatal(err)
	}
	config := oauth2.Config{
		ClientID:     "gokubermanagement-user",
		ClientSecret: "b53f3ca6-0559-4781-a9ee-a35cb0d7a64a",
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:5000/user/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})
}
