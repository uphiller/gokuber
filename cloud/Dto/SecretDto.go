package Dto

import "github.com/jinzhu/gorm"

type SecretDto struct {
	gorm.Model
	Name       string `json:"name"`
	Type       string `json:"type"`
	Access_id  string `json:"access_id"`
	Secret_key string `json:"secret_key"`
}
