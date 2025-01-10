package httpapp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	mwLogger "github.com/go-dairy/main-app/internal/http/middleware/logger"
	godairyv1 "github.com/go-dairy/main-app/protos/gen/go/dairy"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	log        *slog.Logger
	httpServer *http.Server
	port       int
	cancel     context.CancelFunc
}

func New(
	log *slog.Logger,
	port int,
) *App {

	mux := runtime.NewServeMux()

	return &App{
		log: log,
		httpServer: &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: mux,
		},
		port: port,
	}

}

func (a *App) MustRun(grpc_port int) error {
	const op = "httpapp.MustRun"

	log := a.log.With(slog.String("op", op))

	r := chi.NewRouter()

	ctx, cancel := context.WithCancel(context.Background())
	a.cancel = cancel
	mux := runtime.NewServeMux()

	r.Use(middleware.RequestID)
	r.Use(mwLogger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Mount("/", mux)

	if err := godairyv1.RegisterGodairyHandlerFromEndpoint(ctx, mux, "localhost:"+strconv.Itoa(grpc_port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.httpServer.Handler = mux

	log.Info("HTTP server is running", slog.String("addr", a.httpServer.Addr))

	if err := a.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "httpapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping HTTP server", slog.Int("port", a.port))
	a.cancel()
	a.httpServer.Shutdown(context.Background())
}
