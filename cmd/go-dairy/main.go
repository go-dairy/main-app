package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-dairy/main-app/internal/app"
	"github.com/go-dairy/main-app/internal/config"
)

//Generate protobufs files
//go:generate ../../bin/task protogen

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application dairy-app",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPC.Port),
	)

	storagePath := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.StorageUser, cfg.StoragePass, cfg.StorageHost, cfg.StoragePort, cfg.StorageDB)

	log.Info("storage connection", slog.String("path", storagePath))

	application := app.New(log, cfg.GRPC.Port, storagePath)

	go application.GRPCServer.MustRun()
	go application.HttpServer.MustRun(cfg.GRPC.Port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signal := <-stop

	log.Info("stopping application", slog.String("signal", signal.String()))

	application.GRPCServer.Stop()

	log.Info("application stoped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
