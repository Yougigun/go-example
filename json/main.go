package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]interface{}{}
	// err:=json.Unmarshal([]byte("{\"aa\":\"xxx\"}"),&m)
	err:=json.Unmarshal([]byte(""),&m)
	fmt.Println(err,m)
}