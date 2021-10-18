package main

import (
	"fmt"
	"time"

	"github.com/Knetic/govaluate"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("(lang=='en' || lang=='zh') && age == 12 ");
	if err!= nil {
		panic(err)
	}
	parameters := make(map[string]interface{}, 8)
	parameters["lang"] = "en";
	parameters["age"] = 12;
	parameters["gender"] = "male";
	s := time.Now()
	result, err := expression.Evaluate(parameters);
	if err!= nil {
		panic(err)
	}
	fmt.Println(result, time.Since(s).Seconds())
	// result is now set to "false", the bool value.

}