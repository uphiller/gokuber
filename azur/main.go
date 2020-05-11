package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	"net/http"
	"pc/azur/database/models"
	"time"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/health", health)
		v1.GET("/vms", GetVM)
		v1.POST("/info", setInfo)
	}
	//router.Use(cors.Default())
	return router
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:1234@/gokuber_azur?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
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

func GetVM(c *gin.Context) {
	vmClient := getVMClient()
	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	info, _ := vmClient.List(ctx, GroupName())

	c.JSON(http.StatusOK, gin.H{
		"message": info.Values(),
	})
}

func getVMClient() compute.VirtualMachinesClient {
	vmClient := compute.NewVirtualMachinesClient(SubscriptionID())
	a, _ := GetResourceManagementAuthorizer()
	vmClient.Authorizer = a
	vmClient.AddToUserAgent(UserAgent())
	return vmClient
}

type OAuthGrantType int

const (
	OAuthGrantTypeServicePrincipal OAuthGrantType = iota
	OAuthGrantTypeDeviceFlow
)

var (
	armAuthorizer      autorest.Authorizer
	batchAuthorizer    autorest.Authorizer
	graphAuthorizer    autorest.Authorizer
	keyvaultAuthorizer autorest.Authorizer
)

func GetResourceManagementAuthorizer() (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(
		grantType(), Environment().ResourceManagerEndpoint)

	if err == nil {
		// cache
		armAuthorizer = a
	} else {
		// clear cache
		armAuthorizer = nil
	}
	return armAuthorizer, err
}

func getAuthorizerForResource(grantType OAuthGrantType, resource string) (autorest.Authorizer, error) {
	var a autorest.Authorizer
	var err error

	switch grantType {

	case OAuthGrantTypeServicePrincipal:
		oauthConfig, err := adal.NewOAuthConfig(
			Environment().ActiveDirectoryEndpoint, TenantID())
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalToken(
			*oauthConfig, ClientID(), ClientSecret(), resource)
		if err != nil {
			return nil, err
		}
		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeDeviceFlow:
		deviceconfig := auth.NewDeviceFlowConfig(ClientID(), TenantID())
		deviceconfig.Resource = resource
		a, err = deviceconfig.Authorizer()
		if err != nil {
			return nil, err
		}

	default:
		return a, fmt.Errorf("invalid grant type specified")
	}

	return a, err
}

func grantType() OAuthGrantType {
	if UseDeviceFlow() {
		return OAuthGrantTypeDeviceFlow
	}
	return OAuthGrantTypeServicePrincipal
}

type Resource = models.Resource
type Info = models.Info

func setInfo(c *gin.Context) {
	type RequestBody struct {
		SubscriptionId string `json:"subscriptionId" binding:"required"`
		ClientId       string `json:"clientId" binding:"required"`
		ClientSecret   string `json:"clientSecret" binding:"required"`
		TenantId       string `json:"tenantId" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	resource := Info{SubscriptionId: requestBody.SubscriptionId, ClientId: requestBody.ClientId, ClientSecret: requestBody.ClientSecret, TenantId: requestBody.TenantId}
	db.NewRecord(resource)
	db.Create(&resource)
	c.JSON(200, "ok")
}
