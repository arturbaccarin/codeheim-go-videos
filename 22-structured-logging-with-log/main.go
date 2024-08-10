package main

import (
	"log/slog"
	"math/rand"
	"os"
	"runtime"
)

func main() {
	slog.Info("This is an info log!", "version", runtime.Version())
	slog.Error("This is an error log!")
	slog.Warn("This is a warning log!")

	options := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, options).WithAttrs([]slog.Attr{slog.String("app_version", "v1.0.0")}))
	slog.SetDefault(logger)

	slog.Info(
		"This is an info log!",
		slog.String("version", runtime.Version()),
		slog.Int("Random number", rand.Int()),
	)

	slog.Info(
		"Golang rocks!",
		slog.String("version", runtime.Version()),
		slog.Group("OS Info",
			slog.String("OS", runtime.GOOS),
			slog.Int("CPUs", runtime.NumCPU()),
			slog.String("arch", runtime.GOARCH),
		),
	)

	// slog.Info("This is an info log!", "version", runtime.Version())
	// slog.Error("This is an error log!")
	// slog.Warn("This is a warning log!")
	// slog.Debug("This is a debug log!")
}
