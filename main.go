package main

import (
	"container/list"
	"errors"
)

type LRUCache struct {
	MaxSize int64
	Cache   map[string]interface{}
	Recency *list.List
}

// returns a cached element or nil if it wans't found
func (c *LRUCache) Get(key string) interface{} {
	element, contains := c.Cache[key]
	if !contains {
		return nil
	}

	return element
}

// adds or updates cache with a new element
func (c *LRUCache) Set(key string, val interface{}) {
	c.Recency.PushBack(key)
	c.Cache[key] = val
}

// returns the actual cache size
func (c *LRUCache) CurrentSize() int64 {
	return int64(len(c.Cache))
}

// returns a properly initialized LRUCache instance
func NewLRUCache(size int64) (*LRUCache, error) {
	if size < 1 {
		return nil, errors.New("LRUCache: size must a positive integer")
	}

	return &LRUCache{
		MaxSize: size,
		Cache:   make(map[string]interface{}, size),
		Recency: list.New(),
	}, nil
}
