package main

import (
	"flag"
	"fmt"
_	"log"
)

var (
	wordPtr = flag.String("word", "default value", "description")
	intPtr  = flag.Int("num", 1, "an example number")
)

func main() {
	var on bool
	flag.BoolVar(&on, "isOn", false, "already declared variable as flag option")
	var value string
    flag.StringVar(&value,"value","define path","enter a path")
	flag.Parse()
	fmt.Println(*wordPtr, "is the string flag")
	fmt.Println(*intPtr, "is the num flag")
	fmt.Println(on, "is the on  flag")
    fmt.Println(value)
	fmt.Println(flag.Args())
}
