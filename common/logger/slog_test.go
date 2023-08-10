// description:
// @author renshiwei
// Date: 2023/8/10

package logger

import (
	"log/slog"
	"os"
	"testing"
)

func TestSlog(t *testing.T) {
	slog.Info("你好", "计数", 3, "计数2", 6)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	logger.Info("你好", "计数", 3, "计数2", 6)
}

func TestSlogNewJSONHandler(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

	slog.Info("你好", "计数", 3, "计数2", 6)
}

func TestSlogNewTextHandler(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

	slog.Info("你好", "计数", 3, "计数2", 6)
}

func TestSlogWith(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})).With("url", "www.baidu.com")

	slog.SetDefault(logger)

	slog.Info("你好", "计数", 3, "计数2", 6)
}

func TestSlogGroup(*testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	slog.Info("你好", "计数", 3, "计数2", 6, slog.Group("request", "method", "GET", "url", "www.baidu.com"))
}
