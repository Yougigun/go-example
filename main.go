package main

import (
	// "encoding/json"
	"fmt"
)

// "fmt"
// "regexp"

func main() {
	a := A{}
	fmt.Println(a.Add(1))
}

type A struct {
	B
}

type B struct {
}

func (b B) Add(x int) (y int) {
	return x + 1
}
func (b A) Add(x int) (y int) {
	return x + 100
}
