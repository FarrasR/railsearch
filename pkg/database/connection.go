package database

import (
	"fmt"
	"railsearch/pkg/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConn *gorm.DB
var once sync.Once

func InitDB(conf config.DatabaseConfig) *gorm.DB {
	once.Do(func() {
		var err error
		dbConn, err = gorm.Open(mysql.Open(getDSN(conf)), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			panic(err)
		}
	})
	return dbConn
}

func GetConn() *gorm.DB {
	return dbConn
}

func getDSN(conf config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DatabaseName)
}
