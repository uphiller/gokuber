package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"pc/cloud/Config"
)

func SetSecret(b *Secret) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetSecrets(b *[]Secret, user_id string) (err error) {
	if err := Config.DB.Where("user_id = ?", user_id).First(b).Error; err != nil {
		return err
	}
	return nil
}
