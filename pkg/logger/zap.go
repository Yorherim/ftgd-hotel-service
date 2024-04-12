package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func NewZapSugarLogger() *zap.SugaredLogger {
	logger := createLogger()
	defer logger.Sync()

	return logger.Sugar()
}

func createLogger() *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	infoFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/info.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})
	warnFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/warn.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})
	errFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/err.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})

	customWarnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	})
	infoLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	errLevel := zap.NewAtomicLevelAt(zap.ErrorLevel)

	encodeTimeSetting := zapcore.TimeEncoderOfLayout(time.DateTime)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = encodeTimeSetting

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	developmentCfg.EncodeTime = encodeTimeSetting

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, infoLevel),
		zapcore.NewCore(fileEncoder, infoFile, infoLevel),
		zapcore.NewCore(fileEncoder, warnFile, customWarnLevel),
		zapcore.NewCore(fileEncoder, errFile, errLevel),
	)

	return zap.New(core, zap.AddCaller())
}
