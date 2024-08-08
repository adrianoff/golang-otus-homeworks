package memorystorage

import (
	"context"
	"testing"

	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/internalerrors"
	"github.com/adrianoff/golang-otus-homeworks/hw12_13_14_15_calendar/internal/storage/entities"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()
	storage := New()
	t.Run("Create event with no errors", func(t *testing.T) {
		event := entities.Event{
			ID:    0,
			Title: "Some title",
		}

		newID, err := storage.Create(ctx, event)
		require.NoError(t, err)

		err = storage.Delete(ctx, newID)
		require.NoError(t, err)
	})

	t.Run("Update event with no errors", func(t *testing.T) {
		event := entities.Event{
			ID:    0,
			Title: "Some title",
		}

		newID, err := storage.Create(ctx, event)
		require.NoError(t, err)

		updateEvent := entities.Event{
			ID:    0,
			Title: "Update Title",
		}
		err = storage.Update(ctx, newID, updateEvent)
		require.NoError(t, err)
	})

	t.Run("Delete event with no errors", func(t *testing.T) {
		event := entities.Event{
			ID:    0,
			Title: "Some title",
		}

		newID, err := storage.Create(ctx, event)
		require.NoError(t, err)

		err = storage.Delete(ctx, newID)
		require.NoError(t, err)
	})

	t.Run("Delete/update event with errors", func(t *testing.T) {
		event := entities.Event{
			ID:    0,
			Title: "Some title",
		}

		newID, err := storage.Create(ctx, event)
		require.NoError(t, err)

		err = storage.Delete(ctx, newID)
		require.NoError(t, err)

		err = storage.Delete(ctx, newID)
		require.ErrorIs(t, internalerrors.ErrNotExistID, err)

		err = storage.Update(ctx, newID, event)
		require.ErrorIs(t, internalerrors.ErrNotExistID, err)
	})
}
