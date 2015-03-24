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

	/*var str string
	fmt.Scanf("%s", &str)
	fmt.Printf("Got %s\n", str)
	randomize(r, str)
	*/

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
     header := rawCSVdata[0]
     asIs := make([]bool, len(header))
     for i, field := range header {
     	var choice string
     	fmt.Printf("Keep field=%s as is?", field)
     	fmt.Scanf("%s", &choice)
     	fmt.Printf("Choice=%v\n", choice)
     	if choice == "y" {
     		asIs[i] = true
     	}
     }
     fmt.Println(asIs)
     // sanity check, display to standard output
     for _, each := range rawCSVdata[1:] {
             for i, val := range each {
             	var newval string
             	if asIs[i] {
             		newval = val
             	} else {
             		newval = randomize(r, val)
             	}
             	fmt.Printf("%s->%s\t", val, newval)
             }
             fmt.Printf("\n")
     }
}