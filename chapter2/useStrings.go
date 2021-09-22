package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	f("%T %[1]v\n",s.HasPrefix("nikhil","ni"))

}
