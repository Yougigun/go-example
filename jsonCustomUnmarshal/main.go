package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// "fmt"
// "regexp"

func main() {
	re := RegrexTest{}
	err := json.Unmarshal([]byte(`{"regrex":{"hahaha":"fffffff","hahaha":"xxxxx"},"others":"others"}`), &re)
	fmt.Println(err)
	fmt.Println(re)
}

type RegrexTest struct {
	Regrex re `json:"regrex,omitempty"`
	Others string   `json:"others"`
}

type re map[header]reString

type reString string
func (t *reString) UnmarshalJSON(bs []byte) error {
	var tt string
	var err error
	fmt.Println(string(bs))
	err = json.Unmarshal(bs, &tt)
	if err != nil {
		return err
	}
	_, err = regexp.Compile(tt)
	if err != nil {
		return err
	}
	*t = reString(tt)
	return nil
}

type header string

func (h *header) UnmarshalJSON(bs []byte) error {
	var th string
	var err error
	fmt.Println("YYY")
	err = json.Unmarshal(bs, &th)
	if err != nil {
		return err
	}
	fmt.Println(th)
	if th == "hahaha"{
		fmt.Println("I am doing")
	}
	_, err = regexp.Compile(th)
	if err != nil {
		return err
	}
	th = th + "ssss"
	// *h = header(th)
	return nil
}
func (h *re) UnmarshalJSON(bs []byte) error {
	var th map[header]reString
	var err error
	err = json.Unmarshal(bs, &th)
	if err != nil {
		return err
	}
	*h = re(th)
	return nil
}