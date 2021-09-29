package main

import "fmt"

func main() {

	a := 1
	b := complex(0, 0)
	fmt.Printf("a type %T,b type %T\n", a, b)
	fmt.Printf("a value %v,b value %v\n", a, b+1)
}
