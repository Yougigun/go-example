package initA

import (
	"context"
	"fmt"
)

var A = 100

func init(){
	fmt.Println("context in main",context.Background())
	fmt.Printf("%p\n",context.Background())
	a := 100
	fmt.Printf("%p\n",&a)
	fmt.Println("in initA")
}