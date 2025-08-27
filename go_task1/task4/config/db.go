package config

import (
	"fmt"
	"log"
	"os"
	"task4/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(config *Config) (*gorm.DB, error) {

	// 构建DSN (数据源名称)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.DBName,
		config.MySQL.Charset,
	)
	// 配置GORM日志级别
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录不存在错误
			Colorful:                  true,        // 彩色输出
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database object: %w", err)
	}
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("数据库连接成功")

	DB = db
	log.Println("数据库连接成功")

	err1 := db.AutoMigrate(&models.Posts{}, &models.Comments{}, &models.Users{})
	if err1 != nil {
		panic("迁移表结构失败")
	}
	return db, nil
}

func GetDB() *gorm.DB {
	return DB
}
