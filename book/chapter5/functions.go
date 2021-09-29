package main

import (
	"errors"
	"fmt"
)

type Any interface{}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Undefined")
	}
	return a / b, nil
}

func sumAndMul(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}

func PrintAnything(v ...interface{}) {
	fmt.Printf("type %T,value %[1]v\n",v)
}

func main() {

	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(2, -2))

	result, err := divide(2, 0)

	_ = result

	if err != nil {
		fmt.Println(err)
	}

	res, err := divide(10, 5)

	if err == nil {
		fmt.Println(res)
	}

	fmt.Println(sumAndMul(2, 5))
	add, mul := sumAndMul(20, 5)
	fmt.Println(add, mul)

	PrintAnything("nikhil",1)

	slice1 := []Any{1,2,3,"thing"}
	fmt.Println(slice1)
	PrintAnything(slice1)


}
