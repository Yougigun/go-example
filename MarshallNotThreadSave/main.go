package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type app struct {
	M map[string]string `json:"m"`
}

type app2 struct {
	M string `json:"m`
}
func main() {

	// Marshall map cause concurrent write
	// str:= app{M:map[string]string{"a":"a"}}
	// bs,_:=json.Marshal(str)
	// fmt.Println(string(bs))
	// cont := &app{}
	// var wg sync.WaitGroup
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go writeMap(cont,bs,&wg)
	// }
	// wg.Wait()

	// Marshall string is fine
	str2:= app2{M:"a"}
	bs2,_:=json.Marshal(str2)
	fmt.Println(string(bs2))
	cont2 := &app2{}
	var wg2 sync.WaitGroup
	wg2.Add(10)
	for i := 0; i < 10; i++ {
		go writeMap2(cont2,bs2,&wg2)
	}
	wg2.Wait()
}

func writeMap (c *app,bs []byte,wg *sync.WaitGroup) {
	defer wg.Done()
	defer func ()  {
		if r:=recover();r!=nil{
			fmt.Print("recover",r)
		}
	}()
	// json := jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < 10000; i++ {
		err := json.Unmarshal(bs,c)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func writeMap2 (c *app2,bs []byte,wg *sync.WaitGroup) {
	defer wg.Done()
	defer func ()  {
		if r:=recover();r!=nil{
			fmt.Print("recover",r)
		}
	}()
	// json := jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < 10000; i++ {
		err := json.Unmarshal(bs,c)
		if err != nil {
			fmt.Println(err)
		}
	}
}
