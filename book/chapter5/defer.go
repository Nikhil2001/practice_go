package main

import "fmt"

func Print() {

	defer fmt.Println()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}

func main() {

	defer func(s string) {
		fmt.Println(s)
	}("bye user!")
	defer fmt.Println("end!")
	fmt.Println("begin!")
	Print()

	

}
