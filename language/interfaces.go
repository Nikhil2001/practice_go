package main

import "fmt"

type Product interface {
	Print()
}

type watch struct {
	version int
}

func (w watch) Print() {
	fmt.Println("Apple watch series", w.version)
}

func (a AppleDevice) Print() {
	fmt.Println("AppleDevice watch series", a.watch.version)
}

type AppleDevice struct {
	watch
}

func main() {
	device1 := watch{7}
	device1.Print()

	device2 := AppleDevice{watch{version: 8}}
	device2.Print()
}
