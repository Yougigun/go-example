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
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello "+name, ", I am "+hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed " + fmt.Sprint(hero.Speed))
	return "Running"
}


func main() {
	// With reflect.TypeOf method, we can get a type info of variable
	// custom type
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Printf("Hero's type is %s, kind is %s\n", typeOfHero, typeOfHero.Kind())

	// custom Prt type
	typeOfPtrHero := reflect.TypeOf(&Hero{})
	fmt.Printf("PtrHero's type is %s, kind is %s\n", typeOfPtrHero, typeOfPtrHero.Kind())
	
	// Type.Elem return a type's elem type
	typeOfPtrHeroElem := typeOfPtrHero.Elem()
	fmt.Printf("PtrHero's elem type is %s, kind is %s\n", typeOfPtrHeroElem, typeOfPtrHeroElem.Kind())

	// interface type
	var p Person = &Hero{}
	typeOfPerson := reflect.TypeOf(p)
	fmt.Printf("Interface Hero's type is %s, kind is %s\n", typeOfPerson, typeOfPerson.Kind())
}
