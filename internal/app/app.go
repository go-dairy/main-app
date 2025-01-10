package app

import (
	"log/slog"

	grpcapp "github.com/go-dairy/main-app/internal/app/grpc"
	httpapp "github.com/go-dairy/main-app/internal/app/http"
	"github.com/go-dairy/main-app/internal/services/godairy"
	"github.com/go-dairy/main-app/internal/storage/pgsql"
)

type App struct {
	GRPCServer *grpcapp.App
	HttpServer *httpapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {

	storage, err := pgsql.New(storagePath)
	if err != nil {
		panic(err)
	}

	godaryService := godairy.New(log, storage, storage)

	grpcApp := grpcapp.New(log, godaryService, grpcPort)

	httpApp := httpapp.New(log, 8080)

	return &App{
		GRPCServer: grpcApp,
		HttpServer: httpApp,
	}
}
