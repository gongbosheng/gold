package common

import (
	"fmt"
	"os"
	"time"

	"gold/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局日志变量
var Log *zap.SugaredLogger

/**
 * 初始化日志
 * filename 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 * 由于zap不具备日志切割功能, 这里使用lumberjack配合
 */
func InitLogger() {
	fmt.Printf("path is %v", config.Conf.Logs.Path)

	var coreArr []zapcore.Core

	// 获取编码器
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "file",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// 日志级别
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zap.ErrorLevel
	})
	// 当yml配置中的等级大于Error时，lowPriority级别日志停止记录
	if config.Conf.Logs.Level >= 2 {
		lowPriority = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return false
		})
	}

	coreArr = append(coreArr, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lowPriority))

	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	Log = logger.Sugar()
	Log.Info("初始化zap日志完成!")
}
