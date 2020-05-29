package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/jinzhu/gorm"
	"net/http"
	"pc/cloud/Config"
	"pc/cloud/Dto"
	"pc/cloud/Models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	r.Use(cors.New(config))
	r.Use(ginsession.New())
	v1 := r.Group("/v1/cloud")
	v1.Use(VerifyToken)
	{
		v1.GET("/clusters", getClusters)
		v1.POST("/cluster", setCluster)
		v1.GET("/secrets", getSecrets)
		v1.POST("/secret", setSecret)
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
	Config.DB, _ = gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/gokuber?charset=utf8&parseTime=True&loc=Local")
	Config.DB.LogMode(true)
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Cluster{})
	Config.DB.AutoMigrate(&Models.Secret{})

	r := setupRouter()
	r.Run(":5001")
}

func getClusters(c *gin.Context) {
	var cluster []Models.Cluster
	err := Models.GetClusters(&cluster)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"clusters": cluster,
		})
	}
}

func setCluster(c *gin.Context) {
	var clusterDto Dto.ClusterDto
	c.BindJSON(&clusterDto)
	var cluster Models.Cluster
	cluster.Name = clusterDto.Name
	cluster.Type = clusterDto.Type
	cluster.Status = "Active"
	err := Models.SetCluster(&cluster)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"clusters": cluster,
		})
	}
}

func getSecrets(c *gin.Context) {
	var secret []Models.Secret
	err := Models.GetSecrets(&secret, fmt.Sprint(c.MustGet("uid")))
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"secrets": secret,
		})
	}
}

func setSecret(c *gin.Context) {
	var secretDto Dto.SecretDto
	c.BindJSON(&secretDto)
	var secret Models.Secret
	secret.Name = secretDto.Name
	secret.Type = secretDto.Type
	secret.Access_id = secretDto.Access_id
	secret.Secret_key = secretDto.Secret_key
	secret.User_id = fmt.Sprint(c.MustGet("uid"))

	err := Models.SetSecret(&secret)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"clusters": secret,
		})
	}
}

//https://github.com/dedidot/gorm-gin
//https://gorm.io/
//https://github.com/tjrexer/react-notifications
