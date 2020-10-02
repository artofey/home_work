package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	// Place your code here
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Place your code here:
	// - capacity
	// - queue
	// - items
}

type cacheItem struct {
	// Place your code here
}

func (lruCache) Set(key string, value interface{}) bool {
	_, _ = key, value
	return true
}

func (lruCache) Get(key string) (interface{}, bool) {
	_ = key
	return nil, true
}

func (lruCache) Clear() {

}

func NewCache(capacity int) Cache {
	return &lruCache{}
}
