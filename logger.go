package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// New creates a new *zap.Logger with specified settings
// output parameter is for log file. leave if empty, if you need std output.
// debugLevel parameter is for enabling debug level logging
func New(output string, debugLevel bool) *zap.Logger {
	// Настраиваем уровень логирования
	logLevel := zapcore.InfoLevel
	if debugLevel {
		logLevel = zapcore.DebugLevel
	}

	// Дефолтный вывод - консоль
	var writeSyncer zapcore.WriteSyncer
	writeSyncer = os.Stdout

	// Создаем конфигурацию для вывода в файл
	if output != "" && output != "standard" && output != "std" {
		file, _ := os.Create(output)
		writeSyncer = zapcore.AddSync(file)
	}

	// Настраиваем форматирование лога с разделением по уровням и датой
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Дата в формате ISO 8601

	// Настраиваем вывод в консоль и файл
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                // Формат JSON
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writeSyncer), // В консоль и файл
		logLevel, // Уровень логов
	)

	// Создаем логгер
	logger := zap.New(core)
	defer logger.Sync() // Для корректного завершения работы логгера

	return logger
}
