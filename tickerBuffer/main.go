package main

import (
	"fmt"
	"time"
)

func main() {
	tr := time.NewTicker(time.Second*2)
	t := <- tr.C
	time.Sleep(time.Second*3)
	fmt.Println(t) //2s
	t = <- tr.C
	// output is 4s. 
	// because ticker channel is one-size buffer. so even doing something 
	// takes time. ticker still can put tick in the channel 
	// but drop the following tick while channel block.
	fmt.Println(t) 
	
}

