package main

import "fmt"
import "strconv"


func main(){

	intValue, err := strconv.Atoi("1")

	if err !=nil {
		fmt.Println(err)
	}

	fmt.Println(intValue)

	strValue := strconv.Itoa(11)
	fmt.Println(strValue+" is eleven")

	var byteExample byte = 161
//	var byteExample2 uint8 = 73
	fmt.Println(string(byteExample))

	for i,v := range "世世世" {
		fmt.Printf("%T,%T\n",i,v)
		fmt.Printf("%v,%v\n",i,v)
	}

	var ex rune = '世'
	fmt.Printf("%v %T",ex,ex)


	
}