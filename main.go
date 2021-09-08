package main

import (
	"fmt"
	"go-example/conditions"
	"strings"
)

func main() {
    // Our condition to check
    s := "($0 > 0.45) AND ($1 == \"ON\" OR ($lang == \"ACTIVE\" AND $3 == true)) AND $3 == false"
    // Parse the condition language and get expression
    p := conditions.NewParser(strings.NewReader(s))
    expr, err := p.Parse()
    if err != nil {
        // ...
		fmt.Println("Parse error")
		fmt.Println(err.Error())
    }

    // Evaluate expression passing data for $vars
    data := map[string]interface{}{"$0": 0.5, "$1": "OFF", "$2": "ACTIVE", "$3": false}
    r, err := conditions.Evaluate(expr, data)
    if err != nil {
        // ...
		fmt.Println("Evaluate Error")
		fmt.Println(err.Error())
    }

    // r is false
    fmt.Println("Evaluation result:", r)
}