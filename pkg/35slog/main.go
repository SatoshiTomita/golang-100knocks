package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	// ログの出力
	slog.Info("Hello, world!")

	// ログの出力（キーと値のペア）
	slog.Info("値を出力", "a", "A", "b", 2)

	// ログの出力（エラー）
	slog.Error("エラーが発生しました")

	// 独自のロガーを作成
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("Hello, world!")

	// ログの出力（グループ化）
	r, _ := http.NewRequest("GET", "localhost", nil)
	slog.Info("finished",
		slog.Group("req",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String())),
		slog.Int("status", http.StatusOK),
		slog.Duration("duration", time.Second))

	// ログの出力（ネスト）
	slog.Info("hello", slog.Int("count", 3))
	slog.Info("hello", slog.Int("count", 3), slog.Bool("ok", true))

	// ログの出力（フォーマット）
	infof := func(format string, args ...any) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}
	infof("Hello, %s!", "world")

}
