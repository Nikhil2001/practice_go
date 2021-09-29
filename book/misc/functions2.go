package main

import (
	"fmt"
	"strconv"
)

func show(n int) {
	fmt.Println("show -->", n)
}

func incWrong(n int) {
	n++ //will not be reflected in main
}

func inc(n int) int {
	n++
	return n
}

func incBy(n int, factor int) int {
	return n * factor
}

func incByStr(n int, factor string) (f int, err error) {
	f, err = strconv.Atoi(factor)
	return n * f, err
}

func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

func limit(n,limit int) int {
	if n > limit {
		return limit
	}
	return n
}


func main() {
	local := 10
	show(local)
	incWrong(local)
	show(local)
	local = inc(local)
	show(local)
	show(incBy(local, 10))

	multiply, err := incByStr(local, "10")
	fmt.Println("show -->", multiply, err)
	multiply, err = incByStr(local, "nikhil")
	fmt.Println("show -->", multiply, err)
	fmt.Println("show -->", sanitize(multiply, err))
	fmt.Println("show -->", limit(5, 4))
	fmt.Println("show -->", limit(5, 20))
}
