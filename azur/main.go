package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/health", health)
		v1.GET("/vms", GetVM)
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

func GetVM(c *gin.Context) {
	vmName := "kube-dev"
	vmClient := getVMClient()
	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	info, _ := vmClient.Get(ctx, GroupName(), vmName, compute.InstanceView)

	c.JSON(http.StatusOK, gin.H{
		"message": info,
	})
	//return vmClient.Get(ctx, GroupName(), vmName, compute.InstanceView)
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
