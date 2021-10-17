package main

import "fmt"

type Team struct {
	name string
	captain string
}


func (t Team) Print(){
	fmt.Printf("%-15s %s\n",t.name,t.captain)
}

func Print(t Team){
	fmt.Printf("%-15s %s\n",t.name,t.captain)
}

func (t *Team) ChangeCaptain(){
	t.captain = "SKY"
}


func (t Team) String() string{
	return fmt.Sprintf("we are %s. our captain is %s",t.name,t.captain)
}

func main(){
	team := Team{"MI","Rohit"}
	team.Print()
	Print(team)
	team.ChangeCaptain()
	Print(team)
	fmt.Println(team)

}