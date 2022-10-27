// description:
// @author renshiwei
// Date: 2022/10/5 17:33

package initialize

import (
	"github.com/duktig666/go-common/common/global"
	"github.com/duktig666/go-common/common/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"strings"
)

func InitGorm() {
	db, err := gorm.Open(mysql.Open(global.Config.Db.Dsn), &gorm.Config{
		Logger: gormLog.Default.LogMode(switchGormLog()),
	})
	if err != nil {
		logger.Errorf("连接数据库失败, error=" + err.Error())
	}

	global.Gorm = db
}

// gorm 日志配置转换
func switchGormLog() gormLog.LogLevel {
	var logLevel gormLog.LogLevel
	switch strings.ToLower(global.Config.Log.Level.Gorm) {
	case "silent":
		logLevel = gormLog.Silent
	case "error":
		logLevel = gormLog.Error
	case "warn":
		logLevel = gormLog.Warn
	case "info":
		logLevel = gormLog.Info
	default:
		logLevel = gormLog.Silent
	}
	return logLevel
}
