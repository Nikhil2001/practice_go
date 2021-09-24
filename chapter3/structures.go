package main

import "fmt"

type Entry struct {
	Name    string
	SurName string
	Year    int
}

func newEntry() *Entry {
	return new(Entry) // &Entry{}
}


func zeroEntry() Entry {
	return Entry{}
}

func main() {

	entry1 := newEntry()
	fmt.Printf("%#v\n",entry1)
	fmt.Printf("%#v\n",*entry1)

	entry2 := zeroEntry()
	fmt.Printf("%#v\n",entry2)

	fmt.Println(*entry1 == entry2)

	entry3 := Entry{"mara","nikhil",1994}
	fmt.Printf("%#v\n",entry3)

	entry4 := Entry{Year : 1994}
	fmt.Printf("%#v\n",entry4)


}

/*OUTPUT :
&main.Entry{Name:"", SurName:"", Year:0}
main.Entry{Name:"", SurName:"", Year:0}
main.Entry{Name:"", SurName:"", Year:0}
true
main.Entry{Name:"mara", SurName:"nikhil", Year:1994}
main.Entry{Name:"", SurName:"", Year:1994}
*/