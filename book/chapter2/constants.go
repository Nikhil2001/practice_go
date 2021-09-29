package main

import "fmt"

const (
	one = 1 << iota //constant generator
	two 
	four 
	_
	_
    thirtyTwo
)

func main(){

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(four)
	fmt.Println(thirtyTwo)


	const num = 1
	fmt.Printf("%v %[1]T\n",num)
	const numFloat = 1.2
	fmt.Printf("%v %[1]T\n",numFloat)

}