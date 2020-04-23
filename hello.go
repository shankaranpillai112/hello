package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/shankaranpillai112/hello/morestrings"
	"github.com/shankaranpillai112/hello/api"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	api.api.run_API()
}
