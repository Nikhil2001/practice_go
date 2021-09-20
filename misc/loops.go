package main

import "fmt"
import "strings"

func main(){

	for i := 0 ; i <10; i++ {
		fmt.Print(i," ")
	}
	fmt.Println()
   
	sum,i := 0,1
	for {

		if i > 10 {
			break
		}
		sum+=i
		fmt.Println(i,"---->",sum)
		i++

	}
    fmt.Println(strings.Repeat("-",20))
	sum,i = 0,1
	for {

		if i > 10 {
			break
		}
        
		if i%2==0 {
			i++
			continue;
		}
		sum+=i
       
		fmt.Println(i,"----->",sum)
		i++
	}

	fmt.Println(strings.Repeat("-",20))
	fmt.Printf("%5s","X")

	for  i:=1; i <= 5 ;i++ {
		fmt.Printf("%5d",i)
	}
	fmt.Println()

	for  i:=1; i <= 5 ;i++ {
		fmt.Printf("%5d",i)
		for  j:=1; j <= 5 ;j++ {
			fmt.Printf("%5d",i*j)
		}
		fmt.Println()
	}


}