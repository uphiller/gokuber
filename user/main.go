package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
	"gopkg.in/resty.v1"
	"net/http"
)

type LoginCommand struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(ginsession.New())
	v1 := r.Group("/v1/auth")
	{
		v1.GET("/health", health)
		v1.POST("/login", login)
	}

	v1.Use(VerifyToken)
	{
		v1.GET("/current", current)
	}

	return r
}

func VerifyToken(c *gin.Context) {
	SecretKey := "-----BEGIN PUBLIC KEY-----\n" +
		"MIICmzCCAYMCBgFySk7zmzANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDDAZtYXN0ZXIwHhcNMjAwNTI1MDUyODA3WhcNMzAwNTI1MDUyOTQ3WjARMQ8wDQYDVQQDDAZtYXN0ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCxKOiLaKeQRyOEq4B2OHxuKcg9bS+L2uvHEEJYTmDulgYXKSlga8elSJ8TBwGDG90ctYSd4xj8qJ/1JzDCvK0yTkuBU4MkRRzJ8A5Gz/+1bNxfOhDr3jX2GCGws7+yNjs/F3eOPz6Va5XXavcdXLACV3HApXbPbWjNmhSk6kpyMd3P0ELpkDMAc6yvpEjl4UIFZl84LkUVXr6P4TS291S/qI3VB2vw1YCPPvPQLD36wqpOwvurUW58rd5Sf6V103dO5HKVfEl7h7UWlv7Ji+tb+xew6mkEn0IrUIjp4CjmzQrMYueRKahO+q1pZ1iXGKmJZ+nPjTfx8nRE92EZSJ+nAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAIMrPVPqNv/SC/RnEy//kfo++PCGOK1keUnJxCOxuivtv/4Iv7/0MdeCPMd1eQBdVX2AI/8KEah+YDcE5JwvRHcPqaJbAxg+cwxRkam4N1ujfwp4xvWBCpqhXUXwJ1bV3goFTEwbd5qL19wy0zyW1otoBjfj3LvzhT4JxVMZIQcpBK3/YSn6SjMkGvVaAn4BnWKjfd2Q6pcO6xpJG+rVHcmGhv6j+NV3yM9CsJUZysnrWOmOcw62f7FyX7SHzi50tcN9gQ+8Rd9z/44B4ON641Nl3dL33l6vKBsUpI3nVCbt0icM66/tiNKzd/VjRDfNrYaFBCzcYs2ef2+J60hZvgo=" +
		"\n-----END PUBLIC KEY-----"

	reqToken := c.GetHeader("Authorization")

	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		fmt.Println(er)
		c.Abort()
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Unauthorized"))
		return
	}

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		c.Abort()
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Unauthorized"))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid := fmt.Sprint(claims["sub"])
		name := fmt.Sprint(claims["preferred_username"])
		c.Set("uid", uid)
		c.Set("name", name)
		c.Next()
	}
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
	var loginCmd LoginCommand
	c.BindJSON(&loginCmd)

	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "password",
			"username":      loginCmd.Name,
			"password":      loginCmd.Password,
			"client_id":     "gokuber",
			"client_secret": "5a91013b-8c1b-4c67-92c2-920ed10d9258",
		}).Post("http://kubernetes.docker.internal:8080/auth/realms/master/protocol/openid-connect/token")
	if err != nil {
		print(err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": err,
		})
	}
	if result["error"] != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func current(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"uid":  c.MustGet("uid"),
		"name": c.MustGet("name"),
	})
}
