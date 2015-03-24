package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
)

func main() {
	fmt.Printf("hello\n")
	csvfile, err := os.Open("my.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
    reader.FieldsPerRecord = -1 // see the Reader struct information below
    rawCSVdata, err := reader.ReadAll()
     if err != nil {
     	log.Fatal(err)
     }
     // sanity check, display to standard output
     for _, each := range rawCSVdata {
             fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
     }
}