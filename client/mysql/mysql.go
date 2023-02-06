package mysql

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go01-airbnb/config"
	"go01-airbnb/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//log.Fatal("Connected mysql db")
}

func GetClient(ctx context.Context) *gorm.DB {
	if util.IsEnableTx(ctx) {
		return util.GetTx(ctx)
	}

	return db.Session(&gorm.Session{})
}
