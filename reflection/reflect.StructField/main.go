package main

import (
	"fmt"
	"reflect"
)

// Type.NumField() int, get the number of fields of struct
// Type.Field(i int) StructField, according to index, get StructField
// Type.FieldByName(name string) (StructField, bool), according to field name to get StructField
// StructField.Name string, name
// StructField.Type Type, filed type
// StrcutField.Tag StructTag, StructTag type
// StructField.Offset uintptr, byte offset
// StructField.Index []int, field index
// StructField.Anonymous bool, is public?
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
func main() {
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Println("Filed by Index")
	for i := 0; i < typeOfHero.NumField(); i++ {
		fmt.Printf("field' name is %s, type is %s, kind is %s, tag is %s\n",
			typeOfHero.Field(i).Name,
			typeOfHero.Field(i).Type,
			typeOfHero.Field(i).Type.Kind(),
			typeOfHero.Field(i).Tag.Get("json"),
		)
	}
	nameField, _ := typeOfHero.FieldByName("Name")
	fmt.Println("Filed by Name")
	fmt.Printf("field' name is %s, type is %s, kind is %s, tag is %s\n",
		nameField.Name,
		nameField.Type,
		nameField.Type.Kind(),
		nameField.Tag.Get("json"),
	)

}
