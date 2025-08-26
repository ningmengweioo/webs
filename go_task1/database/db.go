package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dst ...interface{}) {
	var err error
	//不能用:= ---注意
	//DB, err = gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	DB, err = gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// 执行数据库迁移
	DB.AutoMigrate(dst...)

}
