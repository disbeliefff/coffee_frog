package main

import (
	"log/slog"
	"os"

	prettyslog "github.com/disbeliefff/coffee_frog/internal/lib/logger/handlers/pretty_slog"
)

func main() {
	log := setupLogger()
	slog.SetDefault(log)

	log.Info("Starting application...")
}

func setupLogger() *slog.Logger {

	var log = setupPrettySlog()

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := prettyslog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
