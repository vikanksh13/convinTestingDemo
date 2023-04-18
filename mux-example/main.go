// main.go

package main

import (
	"fmt"
	"os"
)

func main() {
	a := App{}
	a.Initialize()
	fmt.Println("Server started on port 2999")
	a.Run(":2999")
}