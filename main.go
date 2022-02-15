package main

import "fmt"

// "fmt"
// "regexp"

func main() {
    type A interface{}
	var a map[int]interface{} = map[int]interface{}{}
	var aa A = a
	Aint,ok:= aa.(map[string]interface{})
	
	fmt.Println(Aint==nil,ok)
	
}

func test()(int){
	var a interface{}
	b:= a.(int)
	return b
}