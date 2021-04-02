package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_WhenItsInitializedWithSize_ThenShouldRespectTheGivenSize(t *testing.T) {
	expectedSize := int64(100)
	cache, _ := NewLRUCache(expectedSize)
	assert.Equal(t, expectedSize, cache.MaxSize)
}

func TestLRUCache_WhenCacheSizeIsReached_ThenShouldDropExtraElements(t *testing.T) {
	expectedSize := int64(2)
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("one", 1)
	cache.Set("two", 2)
	cache.Set("three", 3)
	assert.Equal(t, expectedSize, cache.CurrentSize())
}

func TestLRUCache_WhenSetsAnElement_ThenItShouldBeAbleToGetThatElement(t *testing.T) {
	expectedSize := int64(2)
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("one", 1)
	assert.Equal(t, 1, cache.Get("one"))
}

func TestLRUCache_WhenGetsAnElementNotInCache_ThenItShouldReturnNil(t *testing.T) {
	expectedSize := int64(2)
	cache, _ := NewLRUCache(expectedSize)
	assert.Equal(t, nil, cache.Get("one"))
}

func TestLRUCache_WhenMultipleElementsAreSetInCache_ThenItShouldStoreThemInLifoOrder(t *testing.T) {
	expectedSize := int64(2)
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("c", 3)
	cache.Set("b", 2)
	cache.Set("a", 1)
	assert.Equal(t, []interface{}{1, 2, 3}, cache.valuesToSlice())
}
