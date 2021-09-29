package main

import "fmt"

func main(){
	var intSlice []int
	fmt.Println(intSlice,len(intSlice),cap(intSlice))

	intSlice = make([]int,3)
	fmt.Println(intSlice,len(intSlice),cap(intSlice))
	intSlice = make([]int,3,6)
	fmt.Println(intSlice,len(intSlice),cap(intSlice))

	var numSlice = []int{1,2,3} 
	fmt.Println(numSlice,len(numSlice),cap(numSlice))


	for i,v := range numSlice{
		fmt.Println(i,v)
	}

	b := []byte("nikhil reddy")
	fmt.Println(b,len(b),cap(b))
	fmt.Printf("%d\n",b)
	fmt.Printf("%s\n",b)
	fmt.Printf("%s\n",string(b))


	wish := []byte("Hello World 世")
	fmt.Println(wish,len(wish),cap(wish))
	for i,v := range wish{
		fmt.Printf("%v %v %[2]T \n",i,v)
	}

	for i,v := range "Hello World 世"{
		fmt.Printf("%v %v  %[2]T \n",i,v)
	}

}