package main

import (
	"fmt"

	"github.com/pkg/profile"
)

// run 10 MM cache operations (between Set, Get - with hit and miss cases)
func runCacheSamples() {
	cache, _ := NewLRUCache(1000)
	hit := true
	executions := 10000000 // 10MM

	for i := 0; i < executions; i++ {
		if i%2 == 0 {
			cache.Set(fmt.Sprint("key", i), i)
		} else {
			if hit {
				cache.Get(fmt.Sprint("key", i-1)) // there are only even keys on cache
				hit = false
			} else {
				cache.Get(fmt.Sprint("key", i))
				hit = true
			}
		}
	}
}

func profileCPU() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("/tmp")).Stop()
	runCacheSamples()
}

func profileMemory() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath("/tmp")).Stop()
	runCacheSamples()
}

func main() {
	fmt.Println("Profiling LRU Cache via pprof (CPU)")
	profileCPU()

	fmt.Println("Profiling LRU Cache via pprof (Memory)")
	profileMemory()
}
