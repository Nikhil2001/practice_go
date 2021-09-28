package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchString(nameString string) bool {
	byteSlice := []byte(nameString)
	re := regexp.MustCompile(`^[A-Z][a-z]+$`)
	return re.Match(byteSlice)

}
func main() {

	if len(os.Args) !=2 {
		fmt.Println("enter the name")
		return
	}

	nameString := os.Args[1]
	fmt.Println(nameString)
	fmt.Println("Matched :",matchString(nameString))

}
