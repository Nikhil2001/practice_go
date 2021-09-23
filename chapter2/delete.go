package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func deleteElement(a []int, b int) {
	var k int = -1
	for i, v := range a {
		if b == v {
			k = i
			break
		}
	}

	if k != -1 {
		b :=a[:k]
		a = append(b, a[k+1:]...)
	}
	
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("enter the list and  number to be deleted")
		return
	}

	var list []int
	stringList := os.Args[1]
	intStrings := strings.Split(stringList, " ")
	for _, v := range intStrings {
		value, err := strconv.Atoi(v)
		if err == nil {
			list = append(list, value)
		}
	}

	intNum := os.Args[2]
	value, err := strconv.Atoi(intNum)
	if err != nil {
		fmt.Println("enter number in list to be deleted")
		return
	} else {
		deleteElement(list, value)
	}
	fmt.Println("AFTER DELETING list :", list[:len(list)-1])
}
