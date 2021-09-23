package main

import (
	"fmt"
	 "sort"
)

func main() {
	var intSlice = []int{-1, 2, -33, -3, 9, 7, 10}
    sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)
	sort.IntSlice(intSlice).Sort()
	fmt.Println(intSlice)
	fmt.Println(sort.IntSlice(intSlice).Len())
	fmt.Println(sort.IntSlice(intSlice).Search(2))
	sort.IntSlice(intSlice).Swap(2,4)
	fmt.Println(intSlice)
}
