package main

import (
	"fastjson/fuzz"
	"fmt"
)

func main() {
	fmt.Print("Hello")
	fuzz.Fuzz()
	return
}
