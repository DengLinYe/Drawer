package models

import (
	"Drawer/utils"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	// 设置 GORM 日志记录器
	newLogger := logger.New(
		log.New(utils.LogWriter, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到错误
			ParameterizedQueries:      true,        // 参数化查询
			Colorful:                  false,       // 禁用彩色日志
		},
	)

	var err error
	DB, err = gorm.Open(sqlite.Open("drawer.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库初始化失败: " + err.Error())
	}

	DB.AutoMigrate(
		&Transaction{},
		&Category{},
		&Payee{},
		&Account{},
	)
}
