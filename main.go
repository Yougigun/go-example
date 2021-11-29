package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func main() {
    num := 6
    for index :=0; index < num; index++ {
        resp, _:= http.Get("https://www.binance.com")
        _, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("goroutine num = %d\n", runtime.NumGoroutine())
    }
}