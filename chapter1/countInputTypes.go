package main

import "fmt"
import "os"
import "strconv"

func countInputTypes(){
	var total, nrOfInts, nrOfFloats int
	var invalid []string
   for _,v := range os.Args[1:] {
	   _,err := strconv.Atoi(v)
	   if err == nil {
		   total++
           nrOfInts++
		   continue
	   }

	   _,err = strconv.ParseFloat(v,64)
	   if err == nil {
		   total++
           nrOfFloats++
		   continue
	   }

	   invalid = append(invalid,v)
   }
   fmt.Println("total:",total)
   fmt.Println("no of Ints:",nrOfInts)
   fmt.Println("no of Floats:",nrOfFloats)
   fmt.Println("Invalid:",invalid)
}

/*
OUTPUT:
nikhil@nikhil:~/practice_go/chapter1$ go run main.go countInputTypes.go 1 2 w e 3 4 5.0 7.0
Hello World!
total: 6
no of Ints: 4
no of Floats: 2
Invalid: [w e]

*/