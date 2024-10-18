# Logger for Cash Advisor

Here is the logger for Cash Advisor project.

## Installation

```sh
go get -u github.com/zhukovrost/cadv_logger
```

## Usage

```go 
package main

import (
	logger "github.com/zhukovrost/cadv_logger"
	"go.uber.org/zap"
	"time"
)

func main() {
	// Создаем логгер с выводом в файл "app.log" и включенным отладочным режимом
	log := logger.New("app.log", true)

	log.Info("Application started")

	// Пример записи различных уровней логов
	log.Debug("This is a debug message")
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	// Дополнительная информация с использованием поля
	log.Info("Processing request", zap.String("request_id", "1234"))

	// Задержка, чтобы увидеть, что приложение работает
	time.Sleep(2 * time.Second)

	log.Info("Application finished")
}
```
