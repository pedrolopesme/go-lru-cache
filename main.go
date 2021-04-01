package main

type LRUCache struct {
	Size int64
}

// TODO implement
func (c *LRUCache) Get(key string) (val string) {
	return
}

// TODO implement
func (c *LRUCache) Set(key, val string) {
}

func NewLRUCache(size int64) *LRUCache {
	return &LRUCache{Size: size}
}
