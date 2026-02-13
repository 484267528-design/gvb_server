package core

import (
	"fmt"
	"gvb_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.DSN()

	var MysqlLogger logger.Interface //GORM 的日志接口
	if global.Config.System.Env == "debug" {
		//开发环境显示所有sql
		MysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		MysqlLogger = logger.Default.LogMode(logger.Error) //只打印错误的sql
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: MysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
