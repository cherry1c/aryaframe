package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Field = zap.Field

var logger *zap.Logger

func getEncoderConf() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func getWriteSyncers() []zapcore.WriteSyncer {
	var writer []zapcore.WriteSyncer
	logPath := "./log/"
	writer = append(writer, zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath + "log.log", // 日志文件路径
		MaxSize:    128,                 // 每个日志文件保存的大小 单位:M
		MaxAge:     7,                   // 文件最多保存多少天
		MaxBackups: 30,                  // 日志文件最多保存多少个备份
		LocalTime:  true,                // 本地时区
		Compress:   false,               // 是否压缩
	}))
	writer = append(writer, os.Stdout)

	return writer
}

func getLoggerLevel() zap.AtomicLevel {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	return atomicLevel
}

func getLoggerOption() []zap.Option {
	var opts []zap.Option
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))
	opts = append(opts, zap.Development())
	return opts
}

func Init() error {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConf()),
		zapcore.NewMultiWriteSyncer(getWriteSyncers()...),
		getLoggerLevel(),
	)

	// 构造日志
	logger = zap.New(core, getLoggerOption()...)
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
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, fields...)
}
