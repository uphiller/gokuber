package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"pc/gcp/Config"
)

func GetAllCluster(b *[]Cluster) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
