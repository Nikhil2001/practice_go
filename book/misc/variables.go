package main

import "fmt"

var language string //not used but allowed

func main() {
	var hexNum = 0b1001
	fmt.Printf("%T %[1]v\n",hexNum)
	fmt.Println(0x1,0x2,0xa)

	var name string
	fmt.Printf("%s %[1]q\n",name)
	name = "nikhil"
	fmt.Println(name)

	తల := "head in telugu language" //unicode allowed in variable names ?
	fmt.Println(తల)

	var(
	   _name string ="nikhil"
		age int
		weight float64
		swedish = false
	)
	fmt.Println(_name,age,weight,swedish)
	fmt.Printf("%q\n",_name)
	v:= 1
	_ = v

	//also

	var name1 ,age1, weight1, indian = "nikhitha", 27, 50, true
	_ ,_ ,_ ,_ = name1, age1, weight1,indian

	var (
	   height int
	   isOn bool
	   brightness float64
	   s string
	   origin complex64
	   unicodePoint rune
	)

	fmt.Println(height)
	fmt.Println(isOn)
	fmt.Println(brightness)
	fmt.Printf("s (%T) %q\n",s,s)
	fmt.Println(origin)
	fmt.Println(unicodePoint)

	var (
		active bool
		delta int
	)
	fmt.Println(active,delta)

	var firstname,lastname string
	fmt.Printf("%q %q\n",firstname,lastname)

	var isLiquid bool
	_ = isLiquid

	//fmt.Println(program)
	//program := "go"

}


/*OUTPUT

int 9
1 2 10
 ""
nikhil
head in telugu language
nikhil 0 0 false
"nikhil"
0
false
0
s (string) ""
(0+0i)
0
false 0
"" ""
*/
