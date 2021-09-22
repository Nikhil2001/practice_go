package main

import (
        "fmt"
        _ "os"
        "strconv"
        "unicode"
)

func main() {

        //      if len(os.Args) != 2 {
        //              fmt.Println("enter integer value.")
        //              return
        //      }

        fmt.Println(string(49)) //converts int to unicode point, prints 1

        value, _ := strconv.Atoi("333")
        fmt.Println(value)

        value, _ = strconv.Atoi("ab")
        fmt.Println(value)

        value1 := strconv.Itoa(7)
        fmt.Println(value1 + " man")

        fmt.Printf("%T %[1]v %c %c\n", unicode.IsPrint(0x99), 0x50, 0x23)

}