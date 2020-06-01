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

func GetSecrets(b *[]Secret, param *Secret) (err error) {
	if err := Config.DB.Where(param).Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetSecret(b *Secret, param *Secret) (err error) {
	if err := Config.DB.Where(param).First(b).Error; err != nil {
		return err
	}
	return nil
}
