package main

import (
	"fmt"
	"sync"
	"time"
)

func tickerLoop(wg *sync.WaitGroup, quit chan struct{}) {
	defer wg.Done()
	select {
	case <-time.Tick(100 * time.Millisecond):
	case <-quit:
		return
	}
}

func main() {
	const n = 100000
	wg := new(sync.WaitGroup)
	wg.Add(n)
	quit := make(chan struct{})
	fmt.Printf("Making %d tickers\n", n)
	for i := 0; i < n; i++ {
		go tickerLoop(wg, quit)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Signalling close\n")
	close(quit)
	wg.Wait()
	fmt.Printf("Sleeping - examine CPU activity for this process (eg top)\n")
	select {}
}