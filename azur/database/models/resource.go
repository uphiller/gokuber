package models

import (
	"github.com/jinzhu/gorm"
	"github.com/velopert/gin-rest-api-sample/lib/common"
)

type Resource struct {
	gorm.Model
	idx  uint
	name string `sql:"type:text;"`
}

func (p Resource) Serialize() common.JSON {
	return common.JSON{
		"idx":  p.idx,
		"name": p.name,
	}
}
