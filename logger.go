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
	var writeSyncers []zapcore.WriteSyncer

	// Если файл вывода не задан, используем стандартный вывод
	if output == "" || output == "standard" || output == "std" {
		writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
	} else {
		// Создаем файл для логирования
		file, err := os.Create(output)
		if err != nil {
			// Обработка ошибки, если файл не может быть создан
			zap.NewNop() // Используем заглушку, если не можем создать логгер
			return nil
		}
		writeSyncers = append(writeSyncers, zapcore.AddSync(file))
	}

	// Настраиваем форматирование лога с разделением по уровням и датой
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Дата в формате ISO 8601

	// Настраиваем вывод в консоль и файл
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),        // Формат JSON
		zapcore.NewMultiWriteSyncer(writeSyncers...), // Используем массив синхронизаторов
		logLevel, // Уровень логов
	)

	// Создаем логгер
	logger := zap.New(core)
	defer logger.Sync() // Для корректного завершения работы логгера

	return logger
}
