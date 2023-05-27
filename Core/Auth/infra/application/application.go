package application

import (
	"auth/infra/adapters/connection"
	"auth/infra/adapters/postgres"
	"auth/infra/realisations/auth_service"
	"auth/logger"
	"auth/probes"
	"context"
)

type App struct {
	l             logger.Logger
	shutdownFuncs []func(ctx context.Context) error
}

func New(l logger.Logger) *App {
	return &App{
		l: l,
	}
}

func (app *App) Start() error {
	probesApp, _ := probes.New(app.l)
	probesApp.SetStarted()
	err := probesApp.Start()
	if err != nil {
		app.l.Sugar().Fatalf("probes start failed: %w", err)
		return err
	}

	store, err := postgres.New()
	if err != nil {
		app.l.Sugar().Fatalf("store start failed: %w", err)
		return err
	}
	app.shutdownFuncs = append(app.shutdownFuncs, store.Stop)

	auth, err := auth_service.New(store)
	if err != nil {
		app.l.Sugar().Fatalf("creation of business logic failed: %w", err)
		return err
	}

	server, err := connection.New(auth, app.l)
	if err != nil {
		app.l.Sugar().Fatalf("server start failed: %w", err)
		return err
	}
	app.shutdownFuncs = append(app.shutdownFuncs, server.Stop)

	err = server.Start()
	if err != nil {
		app.l.Sugar().Fatalf("server not started: %w", err)
		return err
	}
	probesApp.SetReady()
	return nil
}

func (app *App) Stop(ctx context.Context) error {
	var errStop error
	for i := len(app.shutdownFuncs) - 1; i >= 0; i-- {
		errStop = app.shutdownFuncs[i](ctx)
		if errStop != nil {
			app.l.Sugar().Error(errStop)
		}
	}
	app.l.Info("app stopped")
	return nil
}
