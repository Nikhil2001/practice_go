package main

import (
	"fmt"
	 "os"
	 "time"
)

func main() {

	d,err:= time.Parse("30-01-1994",os.Args[1])
	if err == nil {
		fmt.Println(d)
	}else {
		fmt.Println(err)
	}

}
