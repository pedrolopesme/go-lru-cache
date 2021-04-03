package main

import (
	"container/list"
	"errors"
)

type CacheEntry struct {
	key   string
	value interface{}
}

type LRUCache struct {
	maxSize int
	cache   map[string]*list.Element
	recency *list.List
}

// returns a cached element or nil if it wans't found
func (c *LRUCache) Get(key string) interface{} {
	element, contains := c.cache[key]
	if !contains {
		return nil
	}

	c.cache[key] = c.moveFront(element)
	return element.Value.(CacheEntry).value
}

// adds or updates cache with a new element
func (c *LRUCache) Set(key string, val interface{}) {
	var ele *list.Element
	if oldEle, exists := c.cache[key]; exists {
		ele = c.moveFront(oldEle)
	} else {
		if c.CurrentSize() == c.maxSize {
			c.removeOldest()
		}
		ele = c.recency.PushFront(val)
	}
	ele.Value = CacheEntry{key: key, value: val}
	c.cache[key] = ele
}

// returns the actual cache size
func (c *LRUCache) CurrentSize() int {
	return len(c.cache)
}

// internal func to help on tests. It is not design to
// use outside testing.
func (c *LRUCache) valuesToSlice() (out []interface{}) {
	for element := c.recency.Front(); element != nil; element = element.Next() {
		out = append(out, element.Value.(CacheEntry).value)
	}
	return
}

// internal func to move an elemento to the front of recency Queue
// TODO use c.recency.MoveToFront
func (c *LRUCache) moveFront(ele *list.Element) *list.Element {
	c.recency.Remove(ele)
	return c.recency.PushFront(ele.Value)
}

// internal func that remove the oldest item on the cache
func (c *LRUCache) removeOldest() {
	if c.CurrentSize() == 0 {
		return
	}

	oldest := c.recency.Back()
	delete(c.cache, oldest.Value.(CacheEntry).key)
	c.recency.Remove(oldest)
}

// returns a properly initialized LRUCache instance
func NewLRUCache(size int) (*LRUCache, error) {
	if size < 1 {
		return nil, errors.New("LRUCache: size must a positive integer")
	}

	return &LRUCache{
		maxSize: size,
		cache:   make(map[string]*list.Element, size),
		recency: list.New(),
	}, nil
}
