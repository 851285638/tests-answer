package main

import (
	"crypto/rand"
	"math"
	"math/big"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10000)
	m := make(map[int64]int, 0)
	for i := 0; i < 10000; i++  {
		go func() {
			n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
			println(n.Int64())
			mu.Lock()
			m[n.Int64()] = 1
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	println(len(m))
}