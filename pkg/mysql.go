package pkg

import (
	"context"
	"douya/admin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	mysqlx *gorm.DB
)

func MysqlInit() {
	mysqlCfg := config.DefaultConfig.MysqlCfg
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.Username, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.Db)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open mysql failed")
	}
	mysqlx = d
	sqlDB, _ := mysqlx.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetMysqlClient(ctx context.Context) *gorm.DB {
	return mysqlx.WithContext(ctx)
}
