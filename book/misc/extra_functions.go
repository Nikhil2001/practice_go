package main

import "fmt"

func showN(){
	if N == 0 {
		fmt.Println("nothing to show")
	}
	fmt.Println("The value of N: ",N)
}


func incN(){
	N++
}