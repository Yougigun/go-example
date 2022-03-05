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
	err := json.Unmarshal([]byte(`{"regrex":{"hahaha":"fffffff"},"others":"others"}`), &re)
	fmt.Println(err)
	fmt.Println(re)
}

type RegrexTest struct {
	Regrex map[string]reString `json:"regrex,omitempty"`
	Others string   `json:"others"`
}

type reString string

func (t *reString) UnmarshalJSON(bs []byte) error {
	var tt string
	var err error
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
