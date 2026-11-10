// Package main demonstrates the usage of path constants and the Get function.
package main

import (
	"fmt"
	"net/http"

	"github.com/hu3bi/golangci-lint-bug/hx"
)

func main() {
	fmt.Println("Hello, World!")
	x := Get(ResetPassword)
	fmt.Println(x)

	r := &http.Request{}
	hx.IsRequest(r)
}
