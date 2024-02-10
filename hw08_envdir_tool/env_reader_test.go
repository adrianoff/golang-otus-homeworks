package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Remove = from filename", func(t *testing.T) {
		tempDir, err := os.MkdirTemp(".", "testdir_")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		_, err = os.CreateTemp(tempDir, "test=ignored")
		require.NoError(t, err)

		env, err := ReadDir(tempDir)
		require.NoError(t, err)
		require.Empty(t, env)
	})
}
