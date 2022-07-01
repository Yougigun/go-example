package main

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

func main() {
	sv, err := semver.NewVersion("")
	// sv2, _ := semver.NewVersion("2")
	fmt.Println(sv,err,"xxx yyy"=="xxx yyy")
}