package main

import (
	"fmt"
	"time"
)


func PrintF(k chan int) {
	for range [5]int{} {
		k <- 2
		fmt.Println("I am a printer")
		
	}
   
}

func main() {

	d := make(chan int)
	go PrintF(d)
	<-d
		time.Sleep(10*time.Second)

}
