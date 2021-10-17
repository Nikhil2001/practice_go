package main

import (
	"fmt"
	"time"
)

func PrintNumbers(s *string) {
	for i := 0; i < 5; i++ {
		fmt.Println("PrintNumbers",*s, i)
		if i>2 {
			*s ="d"
		}
	}
}

func main() {
    m:="b"
	PrintNumbers(&m)
	go func() {
		for i := 0; i < 5; i++ {
			//time.Sleep(time.Second)
			fmt.Println("anynomous",m, i)
			if i== 2 {
				m ="k"
			}
		}
	}()

	time.Sleep(10*time.Second)

}
