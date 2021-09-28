package main

import (
	"encoding/csv"
	"fmt"
	"os"
)



func main() {

	if len(os.Args) != 3 {
		fmt.Println("enter input and output path")
		return
	}

	input,output := os.Args[1],os.Args[2]
	f, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll() // returns [][]string

	if err != nil {
		fmt.Println(err)
	}


	csvFile,err1 := os.Create(output)
	if err1 != nil {
		return
	}

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, line := range lines {
		fmt.Println(line)
		err = csvWriter.Write(line)
		if err1 != nil {
			fmt.Println(err)
		}
	}
	csvWriter.Flush()

}
