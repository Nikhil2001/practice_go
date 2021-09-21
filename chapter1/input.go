package main

import "fmt"
import "os"
import _ "flag" // what is this ?

func main() {
	fmt.Print("Enter your name :")
	var name string
    fmt.Scanln(&name)
	fmt.Printf("welcome %s!\n",name) 
	fmt.Println(len(os.Args),os.Args[0])
}