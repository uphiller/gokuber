package Dto

import "github.com/jinzhu/gorm"

type ClusterDto struct {
	gorm.Model
	Name    string `json:"name"`
	Quntity string `json:"quntity"`
}
