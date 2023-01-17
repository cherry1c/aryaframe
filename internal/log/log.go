package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Field = zap.Field

var logger *zap.Logger

const module = "aryaframe"

func Init() error {
	debug := true
	logPath := "./log/"
	hook := lumberjack.Logger{
		Filename:   logPath + "log.log", // 日志文件路径
		MaxSize:    128,                 // 每个日志文件保存的大小 单位:M
		MaxAge:     7,                   // 文件最多保存多少天
		MaxBackups: 30,                  // 日志文件最多保存多少个备份
		Compress:   false,               // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是开发环境，同时在控制台上也输出
	if debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 调用栈 -1
	callerSkip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	field := zap.Fields(zap.String("appName", module))

	// 构造日志
	logger = zap.New(core, caller, development, field, callerSkip)
	logger.Info("log init success.")

	return nil
}

func String(key, val string) Field {
	return zap.String(key, val)
}

func Int(key string, val int) Field {
	return zap.Int(key, val)
}

func ErrorF(err error) Field {
	return zap.Error(err)
}

func Debug(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Fatal(msg, fields...)
}
