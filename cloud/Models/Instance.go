package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"pc/cloud/Config"
)

func SetInstance(b *Instance) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}
