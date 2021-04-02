package main

import (
	"container/list"
	"errors"
)

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

	return element.Value
}

// adds or updates cache with a new element
func (c *LRUCache) Set(key string, val interface{}) {
	var ele *list.Element
	if oldEle, exists := c.cache[key]; exists {
		ele = c.moveFront(oldEle)
	} else {
		ele = c.recency.PushFront(val)
	}
	ele.Value = val
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
		out = append(out, element.Value)
	}
	return
}

// internal func to move an elemento to the front of recency Queue
func (c *LRUCache) moveFront(ele *list.Element) *list.Element {
	c.recency.Remove(ele)
	return c.recency.PushFront(ele.Value)
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
