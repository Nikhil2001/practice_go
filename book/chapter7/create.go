package main

import (
	"fmt"
	"time"
)

func print(x int) {
	fmt.Println("*", x)

}

func main() {

	go func(x int) {
		fmt.Printf("%d\n", x)
	}(10)

	go print(15)

	time.Sleep(time.Second / 100)
	fmt.Println("DONE....")
}
