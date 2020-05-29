package Models

import "github.com/jinzhu/gorm"

type Cluster struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50)"`
	Type   string `gorm:"type:varchar(10)"`
	Status string `gorm:"type:varchar(50)"`
}

func (b *Cluster) TableName() string {
	return "cluster"
}

type Secret struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50)"`
	Type       string `gorm:"type:varchar(10)"`
	User_id    string `gorm:"type:varchar(50)"`
	Access_id  string `gorm:"type:varchar(50)"`
	Secret_key string `gorm:"type:varchar(50)"`
}

func (b *Secret) TableName() string {
	return "secret"
}

type Instance struct {
	gorm.Model
	Instance_id string `gorm:"type:varchar(50)"`
}

func (b *Instance) TableName() string {
	return "instance"
}
