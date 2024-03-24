package storage

import (
	"context"
	"log"

	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/entities"
	memorystorage "github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/sql"
)

type Storage interface {
	Connect(ctx context.Context) error
	Close() error
	Create(ctx context.Context, event entities.Event) (uint64, error)
	Update(ctx context.Context, id uint64, event entities.Event) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context) ([]entities.Event, error)
}

func New(storageType string, connectStr string) Storage {
	switch storageType {
	case "sql":
		return sqlstorage.New(connectStr)
	case "memory":
		return memorystorage.New()
	}

	log.Fatal("Storage misconfiguration")
	return nil
}
