package main

import "fmt"
import o "os" //renaming package
import r "runtime" 

var year = 2021


func name() string {
	return "nikhil"
}

func PrintExamples(){
	fmt.Print("Nikhitha")
	fmt.Print("Dev")
	fmt.Println() //to print next line
	fmt.Print("dev","vijaya",100)
	fmt.Println() //to print next line

}

func PrintfExamples(){
	name1 := "nikhil"
	fmt.Printf("%s,%q\n",name1,name1)
	fmt.Printf("%T\n",name1)
	fmt.Printf("%d %T\n",r.NumCPU(),r.NumCPU())
	fmt.Printf("%s %T\n",name(),name())
	
}

func PrintlnExamples(){
	fmt.Println("Nikhil")
	fmt.Println("dev","vijaya")
	fmt.Println("nikhil")
	fmt.Println("nik\nhil")
	fmt.Println("nik\\nhil") // here \\= \ prints nik\nhil
	fmt.Println("nik\\n\"hil\"") // \" = ", prints nik\n"hil"
	fmt.Println(//"go corona"
	);
	fmt.Println(r.Version())
	fmt.Println(year); fmt.Println("corona"); //gofmt functionality
	fmt.Println(o.Args)
}

func scopeExamples(){
	fmt.Println(year)
	var year = 1995
	fmt.Println(year)
	{
	var year = 1994
	fmt.Println(year)
	}
}


func constExamples(){
	const ok = true //const variable declared and not used but works!
}

func dataTypeExamples(){
	var(
		employee bool = false
		age int = 27
        name string = "nikhil"
	    weight  float64 = 70.5
	)
	fmt.Printf("%T %v %t\n",employee,employee,employee)
	fmt.Printf("%T %v %f\n",weight,weight,weight)
	fmt.Printf("%T %v %s %q\n",name,name,name,name)
	fmt.Printf("I am %[2]v mara, age %[1]v\n",age,name)
	fmt.Printf("%[1]T %[1]v %[1]d\n",age)
	fmt.Printf("precision of float %f, %.0f,%.1f %.2f\n",weight,weight,weight,weight)
    fmt.Printf("value of pi %f, %.2[1]f %[1]v\n",3.14)
/*OUTPUT
bool false false
float64 70.5 70.500000
string nikhil nikhil "nikhil"
I am nikhil mara, age 27
int 27 27
precision of float 70.500000, 70,70.5 70.50
value of pi 3.140000, 3.14 3.14
*/
}

func main() {
 
}

