package main

import (
	"fmt"
	"time"
)

func WriteToChannel(x chan int, i int) {
	x <- i

}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go WriteToChannel(c, i)
	}

	n := 0
	for i := range c {
		n++
		fmt.Println(i)
		if n > 2 {
			//	close(c)
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			close(c)
		}
		fmt.Println(<-c, time.Now().UnixMicro())
	}
}
