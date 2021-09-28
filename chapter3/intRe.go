package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) !=2 {
		fmt.Println("enter the number")
		return
	}

	numString := os.Args[1]
	fmt.Println(numString)

	byteSlice := []byte(numString)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	fmt.Println("Matched :",re.Match(byteSlice))

}
