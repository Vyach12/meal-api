package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vyach12/meal-api/cmd/app"
	"github.com/gomeal/config/pkg/config"
)

func main() {
	ctx := context.Background()

	provider := config.NewProvider(".cfg/values.yaml")

	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	var (
		repos = app.InitRepo(ctx, provider)
		clients = app.InitClients(ctx, provider)
		services = app.InitServices(log, ctx, clients, repos)
		app = app.InitApplication(log, services.MealFetcherService)
	)

	go func () {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Stop()
	log.Info("Application is stopped")

}
