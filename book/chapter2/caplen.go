package main

import "fmt"


func main(){
	a := make([]int ,10)
	fmt.Println(a,len(a),cap(a))
	b := a[1:6]
	fmt.Println(b,len(b),cap(b))

	c := a[1:6:8]
	fmt.Println(c,len(c),cap(c))
	c = append(c,[]int{1,2,3}...)
	fmt.Println(c,len(c),cap(c))
	
}