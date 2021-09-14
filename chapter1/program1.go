package main

import "fmt"

func main() {

	//Write and read data 
	fmt.Println("Hello World!")
	var inputUser1, inputUser2 string
	fmt.Print("Enter input string:")
	fmt.Scanln(&inputUser1,&inputUser2)
	fmt.Println(inputUser1)
	fmt.Println(inputUser2)

}