package main

import (
	"log/slog"
	"os"

	"github.com/disbeliefff/coffee_frog/internal/config"
	prettyslog "github.com/disbeliefff/coffee_frog/internal/lib/logger/handlers/pretty_slog"
)

func main() {
	log := setupLogger()
	slog.SetDefault(log)

	cfg := config.Load("./internal/config/config.yaml")
	log.Info("Config loaded", "config", cfg)

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
