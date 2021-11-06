package core

import (
	"gin-tmp/constant/global"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
}

func (l *Logger) InitLogger() {
	hook := lumberjack.Logger{
		Filename:   global.GtConfig.Zap.FileName,   // 日志文件路径
		MaxSize:    global.GtConfig.Zap.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: global.GtConfig.Zap.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     global.GtConfig.Zap.MaxAge,     // 文件最多保存多少天
		Compress:   global.GtConfig.Zap.Compress,   // 是否压缩
		LocalTime:  global.GtConfig.Zap.LocalTime,  // 备份时间本地化
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		NameKey:        "logger",
		CallerKey:      "line",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,               // 换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder,           // 小写编码器
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.Stamp), // ANSIC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,          // 序列化时间
		EncodeCaller:   zapcore.FullCallerEncoder,               // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// TODO:多服务设置
	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "JkbHelper"))
	// 构造日志
	logger := zap.New(core, caller, development)
	global.GtLogger = logger
}
