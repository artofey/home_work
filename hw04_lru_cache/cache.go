package hw04_lru_cache //nolint:golint,stylecheck

// Key is type of string
type Key string

// Cache is interface of cache.
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, ok := c.items[key]
	if ok {
		item.Value = cacheItem{key, value}
		c.items[key] = item
		c.queue.MoveToFront(item)
	}
	item = c.queue.PushFront(cacheItem{key, value})
	c.items[key] = item

	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
		keyCI := lastItem.Value.(cacheItem).key
		delete(c.items, keyCI)
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, ok
	}
	return nil, ok
}

func (c *lruCache) Clear() {

}

// NewCache make a new chache instance.
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem),
	}
}
