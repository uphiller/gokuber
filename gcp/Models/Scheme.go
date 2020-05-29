package Models

import "github.com/jinzhu/gorm"

type Cluster struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50)"`
	Status string `gorm:"type:varchar(50)"`
}

func (b *Cluster) TableName() string {
	return "cluster"
}
