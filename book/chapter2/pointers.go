package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}


func squareFloat(a float64) float64 {
            return a * a
}


func returnFloatP(a float64) *float64 {
	s := a * a
    return &s
}


func inputFloatP(a *float64) *float64 {
	s := *a  *  *a
    return &s
}


func main(){

	var ex *aStructure
	fmt.Println(ex==nil)

	ex = new(aStructure)
	fmt.Println(ex==nil,ex)
	fmt.Printf("%+v \n",ex)
	fmt.Printf("%#v \n",ex)
	fmt.Printf("%v %v \n",&ex.field1,&ex.field2)

	var ex1 aStructure
	fmt.Println(ex1)

	var num *int
	fmt.Println(num)
	number := 1
	num = &number
	fmt.Println(num)
	fmt.Printf("%d\n",num)

	fmt.Printf("%d\n",returnFloatP(6.1))
	f := 6.1
	fmt.Printf("%d\n",inputFloatP(&f))


}


/*OUTPUT:
true
false &{(0+0i) 0}
&{field1:(0+0i) field2:0} 
&main.aStructure{field1:(0+0i), field2:0} 
0xc000014270 0xc000014280 
{(0+0i) 0}
<nil>
0xc00001c110
824633835792
824633835800
824633835808
*/