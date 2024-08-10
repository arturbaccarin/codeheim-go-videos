package examplewithfile

import (
	"io"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"/tmp/slog_demo.log", // better to get this form an env var
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666)

	if err != nil {
		panic("Error opening the log file")
	}

	defer file.Close()

	multiWriter := io.MultiWriter(file, os.Stderr)

	logger := slog.New(slog.NewJSONHandler(multiWriter,
		&slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))

	slog.SetDefault(logger)

	slog.Info("Go Rocks")
}
