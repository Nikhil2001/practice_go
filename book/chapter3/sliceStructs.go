package main

import "fmt"
import "strconv"

type record struct {
	num  int
	text string
}

func main() {

	s := []record{}

	for i := 0; i < 5; i++ {
		s = append(s, record{i, "text" + strconv.Itoa(i)})
	}

	sum := 0
	for _, v := range s {
		sum += v.num
	}
	fmt.Printf("%+v\n", s)
	fmt.Println(sum)

}
