package main

import "fmt"

type Details struct {
	name string
	age  int
}

type Cricketer struct {
	batter  bool
	details Details
}

func (c Cricketer) PrintDetails() string {
	return fmt.Sprintf("batter:%v,details %+v", c.batter, c.details)
}

type TestCricketer struct {
	testMatches int
	Cricketer   //embedded
}

//func (t TestCricketer) PrintDetails() string{
//	return fmt.Sprintf("matches:%v,details %+v",t.testMatches,t.details)
//}

func main() {

	player := Cricketer{
		details: Details{"Dhoni", 40},
		batter:  true,
	}

	//Print Cricketer
	fmt.Println(player.details.name, player.details.age, player.batter)

	testPlayer := TestCricketer{200, player}

	//Print TestCricketer
	fmt.Println(testPlayer.testMatches)
	fmt.Println(testPlayer.Cricketer.details)
	fmt.Println(testPlayer.batter)
	fmt.Println(testPlayer.PrintDetails())

}
