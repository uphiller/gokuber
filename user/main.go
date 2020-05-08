package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/health", health)
		v1.GET("/user", user)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5001")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func user(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user",
	})
}
