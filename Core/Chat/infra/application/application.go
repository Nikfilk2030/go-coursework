package application

import (
	"chat/infra/adapters/connection"
	"chat/infra/adapters/postgres"
	"chat/infra/realisations/usecases"
	"chat/logger"
	"chat/probes"
)

import (
	"context"
)

type Application struct {
	logger        logger.Logger
	shutdownTasks []func(ctx context.Context) error
}

func CreateApp(logger logger.Logger) *Application {
	return &Application{
		logger: logger,
	}
}

func (app *Application) Begin() error {
	healthCheck, _ := probes.New(app.logger)
	healthCheck.SetStarted()
	err := healthCheck.Start()
	if err != nil {
		app.logger.Sugar().Error("probes failed to start: %w", err)
	}
	storage, err := postgres.New()
	if err != nil {
		app.logger.Sugar().Error("storage failed to start: %w", err)
	}
	app.shutdownTasks = append(app.shutdownTasks, storage.Stop)
	commentsUseCase, err := usecases.New(storage)
	if err != nil {
		app.logger.Error("failed to create business logic")
	}
	server, err := connection.NewChatServer(commentsUseCase, app.logger)
	app.shutdownTasks = append(app.shutdownTasks, server.Stop)
	err = server.Start()
	if err != nil {
		app.logger.Sugar().Error("server failed to start: %w", err)
	}
	healthCheck.SetReady()
	return nil
}

func (app *Application) End(ctx context.Context) error {
	var err error
	for i := len(app.shutdownTasks) - 1; i >= 0; i-- {
		err = app.shutdownTasks[i](ctx)
		if err != nil {
			app.logger.Sugar().Error(err)
		}
	}
	app.logger.Info("application stopped")
	return nil
}
