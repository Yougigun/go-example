package main

import (
	"encoding/json"
	"fmt"
)

type People struct{
	Name string `json:"name"`
}

type Doctor struct{
	People
	Career string `json:"career"`
	Str []string `json:"str"`
	A map[int]string `json:"A"`
	B string `json:"B"`
}

func main() {
	m := map[string]interface{}{}
	// err:=json.Unmarshal([]byte("{\"aa\":\"xxx\"}"),&m)
	err:=json.Unmarshal([]byte(""),&m)
	fmt.Println(err,m)

	d := Doctor{
		People: People{Name: "gary"},
		Career: "doctor",
		A:map[int]string{},
	}
	bs,err := json.Marshal(d)
	fmt.Println(string(bs))
	fmt.Println("err: ",err)
	d2 := Doctor{}
	err = json.Unmarshal(bs,&d2)
	fmt.Printf("%+v, err: %v\n",d2,err)


}