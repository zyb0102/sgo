package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormOpt struct {
	Host     string // 主机ip地址
	Port     string // 主机端口
	DBName   string // 数据库名称
	Username string // 用户名
	Password string // 密码
	Charset  string // 字符集编码
	Loc      string // 时区
	LogLevel string //
}

func NewGorm(o *GormOpt) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		o.Username, o.Password, o.Host, o.Port, o.DBName, o.Charset, "True", o.Loc,
	)
	return gorm.Open(mysql.Open(dsn))
}
