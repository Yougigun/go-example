package main

import (
	"fmt"
	"sync/atomic"
)

func main(){
	type M map[string]string
	var v atomic.Value
	var temp = M{}
	var temp2 = M{}
	temp["sss"]="dddd"
	temp2["sss"]="dddd"
	v.Store(temp)
	v.Store(temp2)
	vv := v.Load().(M)
	fmt.Println(vv)
	temp["xxxx"]="xxxx"
	fmt.Printf("p:%p, value: %v\n",&vv,vv)
	fmt.Printf("p:%p, value: %v\n ",&temp,temp)
	var x *struct{
		A string
	}
	x = nil 
	fmt.Println(x==nil&&x.A=="ssss")
}