package main

import "fmt"

func main() {

	var intslice []int
	fmt.Println(intslice)

	var intslice1 = make([]int, 0)
	fmt.Println(intslice1)

	var intslice3 = make([]int, 1)
	fmt.Println(intslice3, len(intslice3), cap(intslice3))

	intslice3[0] = 0
	intslice3 = append(intslice3, 1, 2)
	fmt.Println(intslice3, len(intslice3), cap(intslice3))

	someSlice := []int{3, 4, 5, 6}
	intslice3 = append(intslice3, someSlice...) //add two slices
	fmt.Println(intslice3, len(intslice3), cap(intslice3))

	var intslice2D = make([][]int, 3)
	fmt.Println(intslice2D, len(intslice2D), cap(intslice2D))

	for i, _ := range intslice2D {
		v := make([]int, i+1)

		for j := 0; j < len(v); j++ {
			v[j] = j + 1
		}
		intslice2D[i] = v
	}

	fmt.Println(intslice2D, len(intslice2D), cap(intslice2D))

	for _, v := range intslice2D {
		for _, k := range v {
			fmt.Print(k, " ")
		}
		fmt.Println()

	}

}
