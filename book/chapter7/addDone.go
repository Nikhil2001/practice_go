package main

import (
	"fmt"
	"sync"
	_ "os"
)

func main() {
	count := 20
	var waitGroup sync.WaitGroup
	i := 0
	for ; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
			fmt.Println(i)
		}(i)
	}
	i = 50
	waitGroup.Wait()
	fmt.Println("\nDone....")

}
