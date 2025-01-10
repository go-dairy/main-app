package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	grpcserver "github.com/go-dairy/main-app/internal/grpc"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	godairyService grpcserver.Godairy,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	grpcserver.Register(gRPCServer, godairyService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}

}

func (a *App) MustRun() error {
	const op = "grpcapp.MustRun"

	log := a.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
