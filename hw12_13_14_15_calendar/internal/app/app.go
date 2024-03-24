package app

import (
	"context"

	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage"
	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/entities"
)

type App struct {
	storage storage.Storage
	logger  logger.Logger
}

type Application interface {
	CreateEvent(ctx context.Context, title string) (uint64, error)
}

func New(logger logger.Logger, storage storage.Storage) Application {
	return &App{
		storage: storage,
		logger:  logger,
	}
}

func (app *App) CreateEvent(ctx context.Context, title string) (uint64, error) {
	return app.storage.Create(ctx, entities.Event{ID: 0, Title: title})
}
