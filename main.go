package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(<-chan int,1)
	var c1 chan<- int
	c1 = c
	fmt.Println(c1)
}