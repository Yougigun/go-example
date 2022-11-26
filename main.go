package main

import (
	"fmt"

)


type tokenType int8
const (
	// Token types
	server tokenType = 1 << iota
	client tokenType = 1 << iota
	portal tokenType = 1 << iota
)
func main() {
	// this means url can be used by server, client and portal token
	var a = testFunc()
	fmt.Println(*a)
}
func testFunc() *int {
	// select the logic from the token type
	var aaa = 1
	return &aaa
} 