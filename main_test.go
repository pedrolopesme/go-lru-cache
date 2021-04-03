package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_WhenItsInitializedWithSize_ThenShouldRespectTheGivenSize(t *testing.T) {
	expectedSize := 100
	cache, _ := NewLRUCache(expectedSize)
	assert.Equal(t, expectedSize, cache.maxSize)
}

func TestLRUCache_WhenSetsAnElement_ThenItShouldBeAbleToGetThatElement(t *testing.T) {
	expectedSize := 2
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("one", 1)
	assert.Equal(t, 1, cache.Get("one"))
}

func TestLRUCache_WhenGetsAnElementNotInCache_ThenItShouldReturnNil(t *testing.T) {
	expectedSize := 2
	cache, _ := NewLRUCache(expectedSize)
	assert.Equal(t, nil, cache.Get("one"))
}

func TestLRUCache_WhenMultipleElementsAreSetInCache_ThenItShouldStoreThemInLifoOrder(t *testing.T) {
	expectedSize := 3
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("c", 3)
	cache.Set("b", 2)
	cache.Set("a", 1)
	assert.Equal(t, []interface{}{1, 2, 3}, cache.valuesToSlice())
}

func TestLRUCache_WhenSetTheSameElementMultipleTimes_ThenItKeepOnlyTheLastVersion(t *testing.T) {
	expectedSize := 3
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("a", 3)
	cache.Set("a", 2)
	cache.Set("a", 1)
	assert.Equal(t, 1, cache.CurrentSize())
	assert.Equal(t, 1, cache.recency.Len())
	assert.Equal(t, 1, cache.Get("a"))
}

func TestLRUCache_WhenGetAnItem_ThenItBeMovedToTheFrontOfRecencyQueue(t *testing.T) {
	expectedSize := 3
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("c", 3)
	cache.Set("b", 2)
	cache.Set("a", 1)
	assert.Equal(t, 1, cache.recency.Front().Value.(CacheEntry).value)

	cache.Get("c")
	assert.Equal(t, 3, cache.recency.Front().Value.(CacheEntry).value)
}

func TestLRUCache_WhenCacheSizeIsReached_ThenShouldDropExtraElements(t *testing.T) {
	expectedSize := 2
	cache, _ := NewLRUCache(expectedSize)
	cache.Set("one", 1)
	cache.Set("two", 2)
	cache.Set("three", 3)
	assert.Equal(t, expectedSize, cache.CurrentSize())
}
