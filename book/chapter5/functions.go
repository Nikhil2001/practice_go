package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Undefined")
	}
	return a / b, nil
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

}
