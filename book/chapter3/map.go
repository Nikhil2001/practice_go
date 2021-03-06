package main

import "fmt"

func main() {

	var aMap map[int]string

	fmt.Printf("%v\n", aMap)
	fmt.Printf("%v\n", aMap == nil)

	//aMap[1]="one"
	//panic: assignment to entry in nil map

	//goroutine 1 [running]:
	//main.main()
	//        /home/nikhil/practice_go/chapter3/map.go:8 +0x39
	//exit status 2

	var bMap = make(map[int]string)
	fmt.Printf("%v\n", bMap)
	fmt.Printf("%v\n", bMap == nil)
	//fmt.Printf("%v\n",bMap == map[int]string{})
	//#command-line-arguments
	//./map.go:14:25:
	//invalid operation: bMap == map[int]string{}
	//(map can only be compared to nil)

	var cMap = map[int]string{}
	fmt.Printf("%v\n", cMap)
	fmt.Printf("%v\n", cMap == nil)

	cMap[1] = "one"
	cMap[2] = "two"
	fmt.Printf("%#v\n", cMap)
	fmt.Printf("%d\n", len(cMap)) //length of a map

	value, isPresent := cMap[4]
	fmt.Printf("%q,%t\n", value, isPresent)

	value, isPresent = cMap[2]
	fmt.Printf("%q,%t\n", value, isPresent)

	for i, v := range cMap {
		fmt.Println(i, "--->", v)
	}
	delete(cMap, 2)

	fmt.Println("after deleting element", cMap)

}

/*
nikhil@nikhil:~/practice_go/chapter3$ go run map.go
map[]
true
map[]
false
map[]
false
map[int]string{1:"one", 2:"two"}
2
"",false
"two",true
1 ---> one
2 ---> two
after deleting element map[1:one]

*/
