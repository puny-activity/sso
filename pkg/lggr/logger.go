package lggr

import (
	"context"
	"log/slog"
	"os"
)

func New(config Config) *slog.Logger {
	return slog.New(&ContextHandler{
		Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLoggerLevel(config.Level()),
		}),
	})
}

func parseLoggerLevel(s string) slog.Level {
	var level slog.Level
	err := level.UnmarshalText([]byte(s))
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "failed to parse logger level",
			slog.String("level", s))
		return slog.LevelInfo
	}
	return level
}
