package main

import "fmt"

func main(){

	for i := 0;i < 10;i++{
		fmt.Print(i," ")
	}
	fmt.Println()

	var i int
	for i < 10 {
		fmt.Print(i," ")
		i++
	}
	fmt.Println()

	i = 0
	for { //infinite loop
		if i == 10 {
			break;
		}
		fmt.Print(i," ")
		i++
	}
    fmt.Println()

	i = 0
	for i < 15 { //infinite loop
		if i%2 != 0 {
			i++
			continue;
		}
		fmt.Print(i," ")
		i++
		
	}
    fmt.Println()
     
}