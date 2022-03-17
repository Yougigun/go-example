package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	module := `
package example.authz

default allow = false

allow {
    input.method == "GET"
    input.path == ["salary", input.subject.user]
}

allow {
    is_admin
}

is_admin {
    input.subject.groups[_] = "admin"
}
`
	// r := rego.New(
	// 	rego.Query("x = data.example.allow"),
	// 	rego.Load([]string{"./example.rego"}, nil))
	// fmt.Println(r)
	ctx := context.Background()
	query, err := rego.New(
		rego.Query("x = data.example.authz.allow"),
		rego.Module("example.rego", module),
	).PrepareForEval(ctx)

	if err != nil {
		// Handle error.
		fmt.Println("Handle error")
	}
	input := map[string]interface{}{
		"method": "GETX",
		"path":   []interface{}{"salary", "bob"},
		"subject": map[string]interface{}{
			"user":   "bob",
			"groups": []interface{}{"sales", "marketing"},
		},
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		// Handle evaluation error.
		fmt.Println("Handle evaluation error.")
	} else if len(results) == 0 {
		// Handle undefined result.
		fmt.Println("Handle undefined result")
	} else if result, ok := results[0].Bindings["x"].(bool); !ok {
		// Handle unexpected result type.
		fmt.Println("Handle unexpected result type.result:",result)
	} else {
		// Handle result/decision.
		fmt.Printf("results:\n")
		fmt.Printf("%+v\n", results) //=> [{Expressions:[true] Bindings:map[x:true]}]
	}
	results, err = query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		// handle error
		fmt.Printf("%+v", results)
	}
	if !results.Allowed() {
		// handle result
		fmt.Printf("results.Allowed():\n")
		fmt.Printf("%+v\n", results)
	}
}
