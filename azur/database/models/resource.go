package models

import (
	"github.com/velopert/gin-rest-api-sample/lib/common"
)

type Resource struct {
	Idx  uint
	Name string `sql:"type:text;"`
}

func (p Resource) Serialize() common.JSON {
	return common.JSON{
		"idx":  p.Idx,
		"name": p.Name,
	}
}
