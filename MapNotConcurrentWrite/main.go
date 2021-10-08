package main

import (
	"sync"
)

type app struct{
	c string
}

var wg sync.WaitGroup
func main() {
	m := map[string]app{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go writeMap(m)
	}
	wg.Wait()
}

func writeMap (c map[string]app) {
	defer wg.Done()
	// json := jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < 10000; i++ {
		// a := c["a"]
		// a.c = "xxx"
		c["a"] = app{}
	}
}
