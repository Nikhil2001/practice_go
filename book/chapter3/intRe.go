package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(numString string) bool {
	byteSlice := []byte(numString)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(byteSlice)

}
func main() {

	if len(os.Args) !=2 {
		fmt.Println("enter the number")
		return
	}

	numString := os.Args[1]
	fmt.Println(numString)
	fmt.Println("Matched :",matchInt(numString))

}
