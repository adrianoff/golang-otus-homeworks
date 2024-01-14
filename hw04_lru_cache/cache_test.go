package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("key1", 101)
		c.Set("key2", 102)
		c.Set("key3", 103)
		c.Set("key4", 104)

		_, key2IsOk := c.Get("key2")
		require.True(t, key2IsOk)
		_, key3IsOk := c.Get("key3")
		require.True(t, key3IsOk)
		_, key4IsOk := c.Get("key4")
		require.True(t, key4IsOk)
		_, key1IsOk := c.Get("key1")
		require.False(t, key1IsOk)
	})

	t.Run("purge old elements logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("key1", 101)
		c.Set("key2", 102)
		c.Set("key3", 103)

		c.Get("key3")
		c.Get("key3")
		c.Get("key1")
		c.Get("key2")
		c.Set("key1", 1001)

		c.Set("key4", 104)

		_, key2IsOk := c.Get("key2")
		require.True(t, key2IsOk)
		_, key1IsOk := c.Get("key1")
		require.True(t, key1IsOk)
		_, key4IsOk := c.Get("key4")
		require.True(t, key4IsOk)
		_, key3IsOk := c.Get("key3")
		require.False(t, key3IsOk)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
