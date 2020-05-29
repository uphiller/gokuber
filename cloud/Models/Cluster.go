package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"pc/cloud/Config"
)

func GetClusters(b *[]Cluster) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func SetCluster(b *Cluster) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}
