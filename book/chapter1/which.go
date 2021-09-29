package main

import "fmt"
import "os"
import  "path/filepath"

func main(){

	if len(os.Args) ==1 {
		fmt.Println("provide arg")
		return
	}

    file := os.Args[1]
	path := os.Getenv("PATH")
	//fmt.Println(path)
	pathSplit := filepath.SplitList(path)
	
	for _,dir := range  pathSplit {

		filePath := filepath.Join(dir,file)
		fileInfo,err:= os.Stat(filePath)
		if err ==nil {
		  mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 { //is executable
					fmt.Println(filePath)
		           return
				}
			}
		}
	}

}