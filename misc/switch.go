package main

import (
	    "fmt"
         "os"    
		 "strings"
		 "time"
	)


func main() {

    if len(os.Args)!=2 {
		fmt.Println("enter city name")
		return
	}	

	switch city := os.Args[1]; city {
	  case "Nirmal":
		fmt.Println(city,"is in India")
	  case "stockholm", "Linkoping":
		fmt.Println(city,"is in sweden")
	  default :
	  fmt.Println(city, "is in ?")
	}

	i:=0
	switch { //no condition, which is true
	case i > 0:
		fmt.Println("its positive")
	case i < 0:
		fmt.Println("its negative")
	default:
		fmt.Println("its zero")
	}

	fmt.Println(strings.Repeat("-",10))
	i = 2 
	switch { //no condition, which is true
	case i > 10:
		fmt.Println("its greater than 10")
		fallthrough
	case i > 0:
		fmt.Println("its not negative")
		fallthrough
	default:
		fmt.Println(i)
	}

	fmt.Println(strings.Repeat("!",10))
	
	switch i = 11; { //no condition, which is true
	case i > 10:
		fmt.Println("its greater than 10")
		fallthrough
	case i > 0:
		fmt.Println("its not negative")
		fallthrough
	default:
		fmt.Println(i)
	}

	fmt.Printf("time is now %d \n",time.Now().Hour()) //TODO play with time
	
	str := "ToLower"
	switch str { //no condition, which is true
	case "ToLower":
		fmt.Println(strings.ToLower("NIkhil"))
		fallthrough
	case "ToUpper":
		fmt.Println(strings.ToUpper("its not negative"))
		fallthrough
	default:
		fmt.Println(strings.Title("marA nikhil"))
	}


}