package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_WhenItsInitializedWithSize_ThenShouldRespectTheGivenSize(t *testing.T) {
	expectedSize := int64(100)
	cache := NewLRUCache(expectedSize)
	assert.Equal(t, expectedSize, cache.Size)
}
