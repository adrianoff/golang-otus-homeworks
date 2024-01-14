package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(item)
		item.Value.(*cacheItem).value = value
	} else {
		if c.capacity <= c.queue.Len() {
			delete(c.items, c.queue.Back().Value.(*cacheItem).key)
			c.queue.Remove(c.queue.Back())
		}

		c.queue.PushFront(&cacheItem{key: key, value: value})
	}
	c.items[key] = c.queue.Front()

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		c.items[key] = c.queue.Front()
		return item.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
