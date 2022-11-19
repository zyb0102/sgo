package initialize

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sgo/global"
	"sgo/utils"
	"time"
)

func InitZap() *zap.Logger {
	// 首先判断日志文件目录存在不存在,不存在则创建
	logDir := global.Config.Log.Directory
	exists, _ := utils.PathExists(logDir)
	if !exists {
		// 不存在就创建日志文件夹
		_ = os.Mkdir(logDir, os.ModePerm)
	}

	zapCore := zapCore()
	zapOptions := options()
	return zap.New(zapCore, zapOptions...)
}

func options() []zap.Option {
	var zapOptions []zap.Option
	// 开启堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	zapOptions = append(zapOptions, caller, development)
	return zapOptions
}

func zapCore() zapcore.Core {
	//
	hook, _ := rotatelogs.New(
		global.Config.Log.Directory+"/%Y-%m-%d.log",                      // 日志文件名
		rotatelogs.WithLinkName(global.Config.Log.Directory+"/last.log"), // 软连文件
		rotatelogs.WithMaxAge(time.Hour*24*30),                           // 日志过期时间
		rotatelogs.WithRotationTime(time.Hour*24),                        //
	)
	// 格式配置
	zapCoreEncoder := zapCoreEncoderConfig()
	// 写入设置
	var writerSyncer zapcore.WriteSyncer

	// debug模式和正常模式下
	if "debug" == global.Config.Log.Level {
		// debug模式下写入到控制台和日志文件
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))

	} else {
		// 生产环境下只写入日志文件
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook))
	}
	// 日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapLevel(global.Config.Log.Level))
	return zapcore.NewCore(zapCoreEncoder, writerSyncer, atomicLevel)
}

// 日志级别
func zapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	case "warn":
		return zap.WarnLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}

func zapCoreEncoderConfig() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "file",
		FunctionKey:      "function",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-05-04 15:02:01 .000 Z0700"),
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "",
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}
