package main

import "fmt"
import o "os" //renaming package
import r "runtime"
import "expackage" 

var year = 2021


func name() string {
	return "nikhil"
}
func main() {
	fmt.Println("Nikhil")
	fmt.Print("Nikhitha")
	fmt.Print("Dev")
	fmt.Println()
	fmt.Println("dev","vijaya")
	fmt.Print("dev","vijaya",100)
	fmt.Println()
	fmt.Println(o.Args)
	const ok = true //const variable declared and not used but works!
	fmt.Println(year)
	var year = 1995
	fmt.Println(year)

	{
	var year = 1994
	fmt.Println(year)
	}
	
	fmt.Println(year); fmt.Println("corona");
    fmt.Printf("%d %T\n",r.NumCPU(),r.NumCPU())

	fmt.Printf("%s %T\n",name(),name())
    fmt.Println(//"go corona"
	);
	fmt.Println(r.Version())
	fmt.Println(expackage.Name())
}

