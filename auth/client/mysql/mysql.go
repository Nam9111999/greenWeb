package mysql

import (
	"context"
	"fmt"
	"green.env.com/auth/config"
	"green.env.com/auth/util"
	"time"

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

	db, err = gorm.Open(mysql.New(mysql.Config{DSN: connectionString}), &gorm.Config{})
	if err != nil {
		util.GetLogger().Fatal(err.Error())
	}

	if cfg.Stage == "LOCAL" {
		db = db.Debug()
	}

	rawDB, _ := db.DB()
	// rawDB.SetConnMaxIdleTime(time.Hour)
	rawDB.SetMaxIdleConns(cfg.MySQL.DBMaxIdleConns)
	rawDB.SetMaxOpenConns(cfg.MySQL.DBMaxOpenConns)
	rawDB.SetConnMaxLifetime(time.Minute * 5)

	err = rawDB.Ping()
	if err != nil {
		util.GetLogger().Fatal(err.Error())
	}

	util.GetLogger().Info("Connected mysql db")
}

func GetClient(ctx context.Context) *gorm.DB {
	if util.IsEnableTx(ctx) {
		return util.GetTx(ctx)
	}

	return db.Session(&gorm.Session{})
}
