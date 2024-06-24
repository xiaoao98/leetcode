package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
 iterations = 1000000
 threads    = 10
)

type (
 data struct {
  key int
  val int
 }
)

var (
 cache     = make(map[int]int, 100)
 syncCache sync.Map
 mutex     sync.Mutex
 ch        = make(chan data)
)

func measure(name string, f func()) {
 fmt.Println("Start measuring", name, "...")
 start := time.Now()
 f()
 taken := time.Since(start)
 fmt.Printf("Finished measuring %s, time taken: %v\n\n", name, taken)
}

func exec(meth0d func(i int)) {
 wg := new(sync.WaitGroup)
 wg.Add(threads)
 for t := 0; t < threads; t++ {
  go func() {
   for i := 0; i < iterations; i++ {
    meth0d(i)
   }
   wg.Done()
  }()
 }
 wg.Wait()
}
func main() {
 measure("Mutex", func() {
  exec(func(i int) {
   mutex.Lock()
   cache[i%100000] += 1
	// cache[i%10] += 1
   mutex.Unlock()
  })
 })
 measure("sync.Map", func() {
  exec(func(i int) {
   elem, _ := syncCache.LoadOrStore(i, new(int32))
   atomic.AddInt32(elem.(*int32), 1)
  })
 })
 measure("Channels", func() {
  go func() {
   for x := range ch {
    cache[x.key] += x.val
   }
  }()
  exec(func(i int) {
   ch <- data{i, 1}
  })
 })
}