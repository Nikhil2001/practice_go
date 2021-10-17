package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	start := time.Now()

	go func() {
		sig := <-sigs
		if sig == syscall.SIGINT {
			duration := time.Since(start)
			fmt.Println("duration is",duration)
			fmt.Println("exiting.....",sig)
			os.Exit(0)
		}
	}()

	for {
		fmt.Println("waiting....","ProcesdId",os.Getpid())
		time.Sleep(1 * time.Second)
	}

}
