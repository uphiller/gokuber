package main

import (
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/user")
	{
		v1.GET("/health", health)
		v1.GET("/login", login)
		v1.GET("/callback", user)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func login(c *gin.Context) {

	provider, err := oidc.NewProvider(c, "http://localhost:8080/auth/realms/master/account/")
	if err != nil {
		// handle error
	}
	oauth2Config := oauth2.Config{
		ClientID:     "Gokubermanagement-user",
		ClientSecret: "b53f3ca6-0559-4781-a9ee-a35cb0d7a64a",
		RedirectURL:  "http://localhost:5000/user/callback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	c.Redirect(http.StatusMovedPermanently, oauth2Config.AuthCodeURL("abc"))
}

func user(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user",
	})
}
