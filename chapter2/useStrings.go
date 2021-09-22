package main

import (
	"fmt"
	s "strings"
	 "unicode"
)

var f = fmt.Printf

func main() {
	f("%T %[1]v\n",s.HasPrefix("nikhil","ni"))
	f("%T %[1]v\n",s.HasSuffix("nikhil","hil"))
	f("%T %[1]v\n",s.Index("nikhil","h"))
	f("%T %[1]v\n",s.Index("nikhil","p"))
	f("%q \n",s.Split("nikhil","i"))
	f("%q \n",s.Replace("nikhil","i","u",1))
	f("%q \n",s.Replace("nik hil"," ","_",-1))
	f("%q \n",s.Replace("  "," ","u",2))
	f("%v \n",s.EqualFold("mara","MARa"))//case insensitive compare
	f("%q \n",s.Fields("mara nikhil reddy"))
	f("%q \n",s.SplitAfter("nikhihm","h"))

	trimNonIntegers := func(c rune) bool {		
		if unicode.IsDigit(c){
		   return false
		}
		return true
	}

f("%q \n",s.TrimFunc("nik222hihm",trimNonIntegers))
}


/*
nikhil@nikhil:~/practice_go/chapter2$ go doc strings.TrimFunc
package strings // import "strings"

func TrimFunc(s string, f func(rune) bool) string
    TrimFunc returns a slice of the string s with all leading and trailing
    Unicode code points c satisfying f(c) removed.

nikhil@nikhil:~/practice_go/chapter2$ go run useStrings.go 
bool true
bool true
int 3
int -1
["n" "kh" "l"] 
"nukhil" 
"nik_hil" 
"uu" 
true 
["mara" "nikhil" "reddy"] 
["nikh" "ih" "m"] 
"222" 
*/
