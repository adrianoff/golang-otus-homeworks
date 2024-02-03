package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

const dataPath = "testdata"

var srcPath = filepath.Join(dataPath, "input.txt")

func TestCopy(t *testing.T) {
	params := []struct {
		offset int64
		limit  int64
	}{
		{offset: 0, limit: 0},
		{offset: 0, limit: 10},
		{offset: 0, limit: 1000},
		{offset: 0, limit: 10000},
		{offset: 100, limit: 1000},
		{offset: 6000, limit: 1000},
	}

	for _, param := range params {
		t.Run(fmt.Sprintf("%v_%v", offset, limit), func(t *testing.T) {
			dst, err := os.CreateTemp(os.TempDir(), "go-copy")
			require.NoError(t, err)
			defer dst.Close()
			err = Copy(srcPath, dst.Name(), param.offset, param.limit)
			require.NoError(t, err)
			expPath := filepath.Join(dataPath, fmt.Sprintf("out_offset%d_limit%d.txt", param.offset, param.limit))
			fmt.Println(expPath)
			expContent, err := os.ReadFile(expPath)
			require.NoError(t, err)
			dstContent, err := os.ReadFile(dst.Name())
			require.NoError(t, err)
			require.Zero(t, bytes.Compare(expContent, dstContent))
		})
	}
}

func TestErrors(t *testing.T) {
	t.Run("no file", func(t *testing.T) {
		err := Copy("file_not_exists.txt", "destination.txt", 0, 0)
		require.EqualErrorf(t, err, "open source file error: open file_not_exists.txt: no such file or directory", "")
	})

	t.Run("offset and limit should be positive integers", func(t *testing.T) {
		err := Copy(srcPath, "some_destination.txt", -100, -10)
		require.EqualErrorf(t, err, "offset and limit should be positive integers", "")
	})

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", "some_destination.txt", 0, 0)
		require.EqualErrorf(t, err, "unsupported file", "")
	})

	t.Run("same file", func(t *testing.T) {
		src, _ := os.CreateTemp(os.TempDir(), "tmp_source.txt")
		err := Copy(src.Name(), src.Name(), 0, 0)
		require.EqualErrorf(t, err, "source and destination files are the same", "")
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy(srcPath, "destination.txt", 1000000000, 0)
		require.EqualErrorf(t, err, "offset exceeds file size", "")
	})
}
