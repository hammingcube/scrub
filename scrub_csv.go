package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var str string
	fmt.Scanf("%s", &str)
	fmt.Printf("Got %s\n", str)
	randomize(r, str)

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
             fmt.Printf("%s, %s, %d\n", each[0], each[1], len(each))
     }
}