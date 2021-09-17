package main

import "fmt"
import "strconv"


func main(){
    v := "v"
    intValue, err := strconv.Atoi(v)

	if err != nil{
		fmt.Printf("cannot convert %s into int,ERROR : %s\n",v,err.Error())
		fmt.Println(intValue+1, " ")
	}
	
	intValue, _ = strconv.Atoi("99")
	fmt.Println(intValue+1, " ")

	rank := 3
	switch rank {
		case 1 :
			fmt.Println("topper of the batch")
		case 2 , 3 :
			fmt.Print("one of the topper,")
			fallthrough
		default :
			fmt.Printf("your rank is %d\n",rank)
	}
	     
}