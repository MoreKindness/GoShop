package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

	username := "fengguanming"
	password := "BkGfGw3dlR25PMr6"
	host := "mysql.sqlpub.com"
	port := 3306
	dbname := "fengguanming"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
	// 初始化数据库连接
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
