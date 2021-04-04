<h1 align="center">
  <br>
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/df/Go_gopher_app_engine_color.jpg/800px-Go_gopher_app_engine_color.jpg" alt="LRUCache" width="200">
  <br>
  Go Least Recent Usage In-Memory Cache
  <br>
  <br>
</h1>

<h4 align="center"> Simple <a href="https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)"> LRU cache </a> written in Go that allows you to store any type of element with a string-typed key. </h4>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/pedrolopesme/go-lru-cache"> <img src="https://goreportcard.com/badge/github.com/pedrolopesme/go-lru-cache" /></a>
</p>
<br>
 

### Installation

Just go get this:

```bash
$ go get pedrolopesme/go-lru-cache
```

### Usage

Once you have imported `LRUCache` lib like this:

```golang

import (
  ...
  "github.com/pedrolopesme/go-lru-cache"
)
```

You should initiate your cache bucket by defining a size to it:

```golang
func main() {
  cache, err := NewLRUCache(10000) // after 10k items, it gonna drop the least accessed items.  
  if err != nil {
    log.Fatal(err)
  }
}
```

then you'll have both methods `Set` and `Get` available for usage.

#### Set 
```golang
cache.Set("my_key", myObj)
```

#### Get 
```golang
cachedObj := cache.Get("my_key")  // cachedObj will be nil if it wasnt available on cache
```


### Development

This project comes with 3 handy targets on Makefile. They are:


##### test

Run all unit tests on the project


##### bench

Run the benchmarks on the project. The current values on my development machine are:

```bash
❯ make bench
Running benchmark tests
go test -run=XXX -bench=.
goos: darwin
goarch: amd64
pkg: github.com/pedrolopesme/go-lru-cache
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkLRUCache_OnlySettingKeys-4                                 	 2930050	       403.8 ns/op
BenchmarkLRUCache_OnlyGettingKeys-4                                 	 9111232	       132.6 ns/op
BenchmarkLRUCache_HalfSettingHalfGettingKeys_AllHits-4              	 3460792	       334.7 ns/op
BenchmarkLRUCache_HalfSettingHalfGettingKeys_HalfHitsHalfMisses-4   	 4097424	       292.9 ns/op
```

##### profiler

Run CPU and Memory profilers via pprof during a test of 10MM ops on the cache. The current output for both profilers are: [cpu](docs/cpu-profile.pdf) and [memory](docs/mem-profile.pdf). 

