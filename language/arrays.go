package main

import (
	"fmt"
	"strings"
)

func addlines() {
	fmt.Println(strings.Repeat("-", 20))
}

func main() {

	var names [5]string

	names[0] = "Rohit"
	names[1] = "Hardik"
	names[2] = "Pollard"
	names[3] = "Bumrah"
	names[4] = "sky"

	for i, v := range names {
		fmt.Println(i, "->", v)
	}

	addlines()
	//access without copies
	for i := range names {
		fmt.Println(i, "->", names[i])
	}

	addlines()
	runs := [5]int{0, 2, 13, 24, 50}
	for i := 0; i < len(runs); i++ {
		fmt.Printf("Runs scored by %s : %d\n", names[i], runs[i])
	}

	addlines()
	var playerAges [5]int
	fmt.Println("players in the squad: ", playerAges) //zero values

	var playerNames [5]string
	playerNames = names // can assign to same types 
	fmt.Println("players batted: ", playerNames)

}
