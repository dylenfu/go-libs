package zap

import (
	"go.uber.org/zap"
	"log"
	"time"
)

func SimpleZapLogger() {
	logger, err:= zap.NewProduction()
	if err != nil {
		log.Println("zap\t-", "new logger error", err.Error())
	}

	defer logger.Sync()
	sugar := logger.Sugar()

	url := "www.baidu.com"
	sugar.Infow("failed to fetch URL",
	// Structured context as loosely typed key-value pairs.
	"url", url,
	"attempts", 3,
	"backoff",time.Second)
	sugar.Infof("Failed to fetch URL: %s", url)
}
