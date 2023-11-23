package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	// ロガーを作成
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// ログ出力
	url := "https://example.com/"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// Dict型で渡すことも可能
	logger.Info("failed to fetch URL",
		zap.Dict("request",
			zap.String("url", url),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		),
	)
}
