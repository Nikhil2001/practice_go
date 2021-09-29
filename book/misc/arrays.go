package main

import "fmt"
import "strings"
import _ "os"
import _ "math/rand"

func main(){


	var screenNo [5]int
	fmt.Printf("%#v,%[1]T\n",screenNo)

	var movies [5]string
	movies[0] = "black widow"
	movies[1] = "spider man"
	movies[2] = "fast and furious"
	movies[3] = "RRR"
	movies[4] = "pushpa"

	fmt.Printf("%#v\n",movies)
	fmt.Printf("%q\n",movies)

	var arr byte = 0x10//uint8
	fmt.Printf("%#v\n",arr) //0x10

	var num =[...]int{1,2,3}
	fmt.Println(num)
	fmt.Println(strings.Repeat("-",20))

    // if len(os.Args) < 2 {
    //  fmt.Println("Enter User Name") 
    //   return
    //  }

    // feelings := [6]string{"happy","sad","awesome","good","terrible","bad"}

   //fmt.Println(os.Args[0]+" feel "+feelings[rand.Intn(len(feelings))])


    var score = [5]int{1,2,3,4,5}
	score1 := score
	fmt.Println(score,score1)
	fmt.Printf("%#v %#v\n",score,score1)

	colors :=[...][2]string{{"blue","red"},{"white","black"}}
	fmt.Println(colors)


	var songs [5]string
	const TWO = 2
	songs = [5]string{1:"shape of you",TWO:"cheap thrills","levitating"}
	fmt.Printf("%#v\n",songs)
//OUTPUT:[5]string{"", "shape of you", "cheap thrills", "levitating", ""}





 }