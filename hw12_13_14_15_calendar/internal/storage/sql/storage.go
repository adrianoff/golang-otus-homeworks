package sqlstorage

import (
	"context"
	"log"
	"time"

	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/entities"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	connectStr string
	connection *sqlx.DB
}

func New(connectStr string) *Storage {
	return &Storage{
		connectStr: connectStr,
	}
}

func (storage *Storage) Connect(ctx context.Context) (err error) {
	sqlCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	storage.connection, err = sqlx.ConnectContext(sqlCtx, "pgx", storage.connectStr)
	if err != nil {
		return err
	}
	return storage.connection.Ping()
}

func (storage *Storage) Close() error {
	if err := storage.connection.DB.Close(); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) Create(ctx context.Context, event entities.Event) (uint64, error) {
	query := "INSERT INTO events (title) values ($1)"
	result, err := storage.connection.ExecContext(ctx, query, event.Title)
	if err != nil {
		return 0, err
	}
	lastInsertedID, err := result.LastInsertId()
	return uint64(lastInsertedID), err
}

func (storage *Storage) Update(ctx context.Context, id uint64, event entities.Event) error {
	query := "UPDATE events SET title = $1 WHERE id = $2"
	_, err := storage.connection.ExecContext(ctx, query, id, event.Title)
	if err != nil {
		return err
	}

	return nil
}

func (storage *Storage) Delete(ctx context.Context, id uint64) error {
	query := "DELETE FROM events WHERE id = $1"
	_, err := storage.connection.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (storage *Storage) List(ctx context.Context) ([]entities.Event, error) {
	query := "SELECT * FROM events"
	rows, err := storage.connection.NamedQueryContext(ctx, query, struct{}{})
	if err != nil {
		return nil, err
	}
	events := make([]entities.Event, 0)

	for rows.Next() {
		var event entities.Event
		err := rows.StructScan(&event)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	return events, nil
}
