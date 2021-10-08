package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	SayHello(string)
	Run() string
}
type Hero struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Speed int    `json:"speed" binding:"required"`
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello "+name, ", I am "+hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed " + fmt.Sprint(hero.Speed))
	return "Running"
}

// Type.Method(int) Method, get Method by index
// Type.MethodByName(string) Method, get Method by name
// Type.NumMethod() int, get num of method
// Method.Name string, name
// Method.Type Type, type
// Method.Func Value, reflect object can invoke
// Method.Index int, index of method
func main() {
	var p Person = &Hero{}
	typeOfPerson := reflect.TypeOf(p)
	// method by index
	fmt.Println("method by index")
	for i := 0; i < typeOfPerson.NumMethod(); i++ {
		fmt.Printf("method is %s, type is %s, kind is %s. \n",
			typeOfPerson.Method(i).Name,
			typeOfPerson.Method(i).Type,
			typeOfPerson.Method(i).Type.Kind(),
		)
	}
	// method by name
	fmt.Println("method by name")
	method, _ := typeOfPerson.MethodByName("Run")
	fmt.Printf("method is %s, type is %s, kind is %s.\n",
		method.Name, method.Type, method.Type.Kind(),
	)
}
