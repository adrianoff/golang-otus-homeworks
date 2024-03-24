package memorystorage

import (
	"context"
	"sync"

	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/internalerrors"
	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/entities"
)

type Storage struct {
	lastInsertedID uint64
	events         []entities.Event
	eventsMap      map[uint64]*entities.Event
	mutex          sync.Mutex
}

func New() *Storage {
	return &Storage{
		lastInsertedID: 0,
		events:         make([]entities.Event, 0),
		eventsMap:      make(map[uint64]*entities.Event, 0),
	}
}

func (storage *Storage) Connect(_ context.Context) error {
	return nil
}

func (storage *Storage) Close() error {
	return nil
}

func (storage *Storage) Create(_ context.Context, event entities.Event) (uint64, error) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.lastInsertedID++
	event.ID = storage.lastInsertedID
	storage.events = append(storage.events, event)
	storage.eventsMap[storage.lastInsertedID] = &event

	return storage.lastInsertedID, nil
}

func (storage *Storage) Update(_ context.Context, id uint64, event entities.Event) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	oldEvent, ok := storage.eventsMap[id]
	if !ok {
		return internalerrors.ErrNotExistID
	}

	oldEvent.Title = event.Title

	return nil
}

func (storage *Storage) Delete(_ context.Context, id uint64) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	if _, ok := storage.eventsMap[id]; !ok {
		return internalerrors.ErrNotExistID
	}
	for i, e := range storage.events {
		if e.ID == id {
			storage.events = append(storage.events[:i], storage.events[i+1:]...)
			break
		}
	}
	delete(storage.eventsMap, id)
	return nil
}

func (storage *Storage) List(_ context.Context) ([]entities.Event, error) {
	return storage.events, nil
}
