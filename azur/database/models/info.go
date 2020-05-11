package models

import (
	"github.com/jinzhu/gorm"
	"github.com/velopert/gin-rest-api-sample/lib/common"
)

type Info struct {
	gorm.Model
	idx            uint
	subscriptionId string `sql:"type:text;"`
	clientID       string `sql:"type:text;"`
	clientSecret   string `sql:"type:text;"`
}

func (p Info) Serialize() common.JSON {
	return common.JSON{
		"idx":            p.idx,
		"subscriptionId": p.subscriptionId,
		"clientID":       p.clientID,
		"clientSecret":   p.clientSecret,
	}
}
