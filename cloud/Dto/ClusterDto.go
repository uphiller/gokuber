package Dto

import "github.com/jinzhu/gorm"

type ClusterDto struct {
	gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"`
	Quntity string `json:"quntity"`
}
