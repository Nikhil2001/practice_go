package main

import (
	"fmt"
	 "os"
	_ "path"
)

type contact struct {
	name string
	surname string
	number string
}

var contactList = []contact{{"mara","nikhil","+5678166033"},{"gummadi","varalakshmi","+5678165533"},{"tuck","Jagadish","+56781995533"}}

func list(){
	fmt.Println(contactList)
}

func search(name string) *contact{
	for _, contact := range contactList {
		if contact.name == name{
			return &contact
		}
	}
	return nil
}
func main() {

	switch os.Args[1] {
	case "list" :
		  list()
	case "search" :
		found := search(os.Args[2])
		if found == nil {
			fmt.Println("No contact found")	
		}else {
			fmt.Println(*found)
		}
	default :
		fmt.Println("Not supported for now, try search or list as below")
	//	bin := path.Base(os.Args[0])
	//	fmt.Printf("%s search | list \n",bin)	
		fmt.Printf("%s search | list \n",os.Args[0])		 
	}

}