package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	target := "hello, 你好嗎？"
	strtarget:= base64.StdEncoding.EncodeToString([]byte(target))
	fmt.Println(strtarget)
	bs,_:=base64.StdEncoding.DecodeString(strtarget)
	fmt.Println(string(bs))
	
}

