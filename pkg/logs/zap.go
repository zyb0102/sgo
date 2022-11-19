package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var logger *zap.SugaredLogger

const storeDir = "logs"

func Init() {
	// 判断文件存放目录是否存在
	exists, err := pathExists(storeDir)
	if err != nil {
		panic(err)
	}
	if !exists {
		// 不存在就创建日志文件夹
		_ = os.Mkdir(storeDir, os.ModePerm)
	}
	// 实现两个判断日志等级的interface
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	infoWriter := getWriter(storeDir + "/info")
	errorWriter := getWriter(storeDir + "/error")
	encoder := zapCoreEncoderConfig()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), debugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	log := zap.New(core, zap.AddCaller())
	logger = log.Sugar()
}

func getWriter(fileName string) io.Writer {
	hook, err := rotatelogs.New(
		fileName+"_%Y%m%d.log",                    // 日志文件名
		rotatelogs.WithLinkName(fileName+".log"),  // 软连文件
		rotatelogs.WithMaxAge(time.Hour*24*30),    // 日志过期时间
		rotatelogs.WithRotationTime(time.Hour*24), // 分割一次日志
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func zapCoreEncoderConfig() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		TimeKey:        "ts",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		CallerKey:      "file",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)

}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
