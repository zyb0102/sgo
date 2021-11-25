package global

import (
	"gin-api/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	Today = time.Now().Format("2006-01-02")
	// Viper  库解析处理配置参数
	Viper *viper.Viper
	// Config 配置参数
	Config config.Config
	// Logger zap日志处理
	Logger *zap.Logger
	// DB 数据库实例
	DB *gorm.DB
	// Redis 缓存
	Redis *redis.Client
)
