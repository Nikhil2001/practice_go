package main

import ("fmt"
        "os"
        "strconv"
)

func minMax() {
	var min,max int64
	for i,v := range os.Args[1:] {
        k,_ := strconv.ParseInt(v,10,0) //same as strconv.Atoi(v)
		if i == 0{
			min,max = k,k
			continue 
		}
		if k < min {
			min = k 
		}
		if k > max {
			max = k 
		}
		fmt.Println("the value of k",k)
	}

	fmt.Println("minimum:",min,"maximum:",max)
}