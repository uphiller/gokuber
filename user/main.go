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
		v1.GET("/callback", callback)
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
	var provider, err = oidc.NewProvider(c, "http://localhost:8080/auth/realms/master")
	if err != nil {
		return
	}
	var oauth2Config = oauth2.Config{
		ClientID:     "gokubermanagement-user",
		ClientSecret: "b53f3ca6-0559-4781-a9ee-a35cb0d7a64a",
		RedirectURL:  "http://localhost:5000/v1/user/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	c.Redirect(http.StatusMovedPermanently, oauth2Config.AuthCodeURL("login"))
}

func callback(c *gin.Context) {
	var provider, err = oidc.NewProvider(c, "http://localhost:8080/auth/realms/master")
	if err != nil {
		return
	}
	var oauth2Config = oauth2.Config{
		ClientID:     "gokubermanagement-user",
		ClientSecret: "b53f3ca6-0559-4781-a9ee-a35cb0d7a64a",
		RedirectURL:  "http://localhost:5000/v1/user/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oauth2Token, err := oauth2Config.Exchange(c, c.Request.URL.Query().Get("code"))
	if err != nil {
		return
	}

	userInfo, err := provider.UserInfo(c, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		return
	}

	resp := struct {
		OAuth2Token *oauth2.Token
		UserInfo    *oidc.UserInfo
	}{oauth2Token, userInfo}
	c.JSON(http.StatusOK, resp)
}
