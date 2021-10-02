package logger

import (
	"fmt"
	"os"

	config "gin-i18n/pkg/conf"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// AutoLogger is a utility struct for logging data in an extremely high performance system.
// We can use both Logger and SugarLog for logging. For more information,
// just visit https://godoc.org/go.uber.org/zap
type AutoLogger struct {
	// Sugar for logging
	*zap.SugaredLogger
	// configuration
	config map[string]interface{}
	// Logger for logging
	Logger *zap.Logger
}

// AtLog is logger
var AtLog *AutoLogger

func init() {
	InitLoggerDefaultDev()
}

// InitLoggerDefault -- init logger default
func InitLoggerDefault(enableDebug bool) {
	// init production encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.MessageKey = "message"
	// init production config
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}
	if enableDebug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	// build logger
	logger, _ := cfg.Build()

	sugarLog := logger.Sugar()
	cfgParams := make(map[string]interface{})
	AtLog = &AutoLogger{sugarLog, cfgParams, logger}
}

// InitLoggerDefaultDev -- init logger dev
func InitLoggerDefaultDev() {
	// init development encoder config
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	// init development config
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig = encoderCfg
	cfg.OutputPaths = []string{"stdout"}
	// build logger
	logger, _ := cfg.Build()

	sugarLog := logger.Sugar()
	cfgParams := make(map[string]interface{})
	AtLog = &AutoLogger{sugarLog, cfgParams, logger}
}

// InitLoggerFile -- init logger write to file
func InitLoggerFile() {
	if _, err := os.Stat("./config/log.toml"); os.IsNotExist(err) {
		viper.Set("LOGGER_OUT_PUT_PATH", "./log/app.log")
		viper.Set("LOGGER_MAX_SIZE_IN_MB", 10)
		viper.Set("LOGGER_MAX_BACKUPS", 10)
		viper.Set("LOGGER_MAX_AGE", 30)
	} else {
		config.ReadConfig("log", "./conf")
	}

	var err error
	AtLog, err = NewLogger(
		viper.GetString("LOGGER_OUT_PUT_PATH"),
		viper.GetInt("LOGGER_MAX_SIZE_IN_MB"),
		viper.GetInt("LOGGER_MAX_BACKUPS"),
		viper.GetInt("LOGGER_MAX_AGE"),
	)

	if err != nil {
		panic(fmt.Sprintf("Cannot create logger with the following error: %s", err))
	}
}

// InitLoggerFileDev -- init logger write to file with development config
func InitLoggerFileDev() {
	if _, err := os.Stat("./config/log.toml"); os.IsNotExist(err) {
		viper.Set("LOGGER_OUT_PUT_PATH", "./log/app.log")
		viper.Set("LOGGER_MAX_SIZE_IN_MB", 10)
		viper.Set("LOGGER_MAX_BACKUPS", 10)
		viper.Set("LOGGER_MAX_AGE", 30)
	} else {
		config.ReadConfig("log", "./config")

	}
	var err error
	AtLog, err = NewLoggerFileDev(
		viper.GetString("LOGGER_OUT_PUT_PATH"),
		viper.GetInt("LOGGER_MAX_SIZE_IN_MB"),
		viper.GetInt("LOGGER_MAX_BACKUPS"),
		viper.GetInt("LOGGER_MAX_AGE"),
	)

	if err != nil {
		panic(fmt.Sprintf("Cannot create logger with the following error: %s", err))
	}
}

// NewLogger create new logger based on file path
func NewLogger(outFilePath string, maxSizeInMB, maxBackups, maxAge int) (*AutoLogger, error) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   outFilePath,
		MaxSize:    maxSizeInMB, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
		Compress:   true,
		LocalTime:  true,
	})

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core)

	sugarLog := logger.Sugar()

	cfgParams := make(map[string]interface{})
	cfgParams["output"] = outFilePath
	cfgParams["maxSize"] = maxSizeInMB
	cfgParams["maxBackup"] = maxBackups
	cfgParams["maxAge"] = maxAge

	return &AutoLogger{sugarLog, cfgParams, logger}, nil
}

// NewLoggerFileDev create new logger based on file path
func NewLoggerFileDev(outFilePath string, maxSizeInMB, maxBackups, maxAge int) (*AutoLogger, error) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   outFilePath,
		MaxSize:    maxSizeInMB, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
		Compress:   true,
		LocalTime:  true,
	})

	cfg := zap.NewDevelopmentEncoderConfig()

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core)

	sugarLog := logger.Sugar()

	cfgParams := make(map[string]interface{})
	cfgParams["output"] = outFilePath
	cfgParams["maxSize"] = maxSizeInMB
	cfgParams["maxBackup"] = maxBackups
	cfgParams["maxAge"] = maxAge

	return &AutoLogger{sugarLog, cfgParams, logger}, nil
}
