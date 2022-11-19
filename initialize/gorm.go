package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sgo/global"
	"time"
)

func InitGorm() *gorm.DB {
	// mysql配置参数
	mysqlConfigParams := global.Config.Mysql
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		mysqlConfigParams.UserName,
		mysqlConfigParams.Password,
		mysqlConfigParams.Host,
		mysqlConfigParams.Port,
		mysqlConfigParams.Database,
		mysqlConfigParams.Charset,
		mysqlConfigParams.ParseTime,
		mysqlConfigParams.Loc,
	)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, _ := gorm.Open(mysql.New(mysqlConfig), gormConfig())
	return db
}

// 先默认日志
func gormConfig() *gorm.Config {
	var config gorm.Config
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond * 100, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,                   // 禁用彩色打印
		},
	)
	config.Logger = newLogger
	return &config
}
