package main

import (
	"fmt"
	"reflect"
)

type a interface{test()}
type aa int
func (z *aa)test(){}
func main() {
	var c aa = 123
	var b a = &c
	val := reflect.ValueOf(b)
	ele := val.Elem() // return ele of pointer, if val is not pointer, panics.
	tp := val.Type()
	kd := val.Kind()
	fmt.Println("interface's dynamic value:",val)
	fmt.Println("interface's dynamic value is nil?:",val.IsNil())
	fmt.Println("interface's dynamic value's pointer's ele:",ele)
	fmt.Println("interface's dynamic value's pointer's ele is int?:",ele.Type())
	fmt.Println("interface's dynamic value's pointer's ele is int?:",ele.Kind() == reflect.Int)
	fmt.Println("interface's dynamic type:",tp)
	fmt.Println("interface's dynamic type's kind:",kd)
	fmt.Println("interface's dynamic type's kind is int?:",kd == reflect.Ptr)

	// 通過 reflect.TypeOf 方法, 我們可以獲取一個變量的類型訊息
}


