package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"pc/gcp/Config"
)

func GetSecrets(b *[]Secret) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func SetSecret(b *Secret) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}
