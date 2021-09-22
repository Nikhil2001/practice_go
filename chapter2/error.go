package main

import (
	"fmt"
	"errors"
	"strconv"
)

func check(a,b int) error {
	if a==0 && b ==0 {
		return errors.New("both a and b are zero")
	}

	if a==0 || b ==0 {
		return fmt.Errorf("either %v  or %v are zero",a,b)
	}
	return nil
}


func main() {

	var err error = check(0,0)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(err.Error() == "both a and b are zero" )
	intValue,err := strconv.Atoi("123")
	fmt.Println(intValue,err)

	_,err = strconv.Atoi("a123")
	fmt.Println(intValue,err)

	err = check(5,0)
	fmt.Println(err.Error() == "either 5  or 0 are zero" )

	fmt.Printf("the data type is %T\n",err)
}