package models

import (
	"github.com/velopert/gin-rest-api-sample/lib/common"
)

type Info struct {
	idx            uint
	SubscriptionId string `sql:"type:text;"`
	ClientId       string `sql:"type:text;"`
	ClientSecret   string `sql:"type:text;"`
	TenantId       string `sql:"type:text;"`
}

func (p Info) Serialize() common.JSON {
	return common.JSON{
		"idx":            p.idx,
		"subscriptionId": p.SubscriptionId,
		"clientID":       p.ClientId,
		"clientSecret":   p.ClientSecret,
		"tenantID":       p.TenantId,
	}
}
