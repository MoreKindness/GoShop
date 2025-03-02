package sql

import (
	"gomall/dal/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

//func Init() *gorm.DB {
//
//	username := "fengguanming"
//	password := "BkGfGw3dlR25PMr6"
//	host := "sql.sqlpub.com"
//	port := 3306
//	dbname := "fengguanming"
//	timeout := "10s"
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
//	// 初始化数据库连接
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//
//	err = dao.InitTables(DB)
//	if err != nil {
//		panic(err)
//	}
//	return DB
//}

func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
