package main

import "fmt"
import "strings"



func isTrue(n *int) bool {
	(*n) += 1
	return false
}

func main() {

	var a,b int = 1,-2
	positive := a > 0
	negative := b < 0
	fmt.Printf(`a(value=%d) is positive ? %t \nb(value=%d) is negative ? %t\n`,a,positive,b,negative)

	c := 1
	var s string = string(c)
	fmt.Println(s) // prints nothing
	//NOTE: > >= are only done between numeric operands
	fmt.Println(strings.Repeat("i",10))

	pi := 3.14
	var num = 4
	fmt.Println(pi > 3 )
	fmt.Println(pi > float64(num))

	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", false && false)
	n:=9
	fmt.Println(true || isTrue(&n) , n ) // short circuited ? isTrue is not called.
	fmt.Println(false || isTrue(&n) , n ) // isTrue is called here.

	var  isOn bool
	fmt.Println(isOn,!isOn,!!isOn,!!!isOn)

	l := false

	if l {
		fmt.Println("true")
	} else {
		fmt.Println("false")

	}

	m :=0
	if m > 0 {
		fmt.Println("positive")
	} else if m <0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}

	fmt.Printf("%g %[1]f \n", 2.333333333*2)

	str :=`     a  b  c  d  e  \nf nikhil`
	fmt.Println(str)
	trim := strings.TrimSpace(str)
    fmt.Println(trim)

}