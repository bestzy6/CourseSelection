package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// InitMysql 初始化mysql链接
func InitMysql() {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	connString := account + ":" + password + "@tcp(" + addrMYSQL + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("mysql lost:", err)
	}
	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	DB = db
}
