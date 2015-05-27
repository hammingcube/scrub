package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))

	/*var str string
	fmt.Scanf("%s", &str)
	fmt.Printf("Got %s\n", str)
	randomize(r, str)
	*/

	csvfile, err := os.Open("Sheet1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	var b bytes.Buffer
	writer := csv.NewWriter(bufio.NewWriter(&b))

	reader.FieldsPerRecord = -1 // see the Reader struct information below
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	header := rawCSVdata[0]
	asIs := make([]bool, len(header))
	/*for i, field := range header {
	  	var choice string
	  	fmt.Printf("Keep field=%s as is?", field)
	  	fmt.Scanf("%s", &choice)
	  	fmt.Printf("Choice=%v\n", choice)
	  	if choice == "y" {
	  		asIs[i] = true
	  	}
	  }
	  for i := 0; i < len(asIs); i++ {
	  	fmt.Printf("%s=%v,", header[i], asIs[i])
	  }
	*/

	//fmt.Printf("Here-->%q", jumble('A'))
	m := make(map[string]func(rune) rune)
	fmt.Println()
	// sanity check, display to standard output
	for _, each := range rawCSVdata[1:] {
		k := each[0]
		if _, ok := m[k]; !ok {
			m[k] = NewJumbler()
		}
		newRecord := make([]string, len(each))
		for i, val := range each {
			var newval string
			if asIs[i] {
				newval = val
			} else {
				jumble := m[k]
				for _, ch := range val {
					newval += string(jumble(ch))
				}
				//newval = jumble(val)randomize(r, val)
			}
			newRecord[i] = newval
			//fmt.Printf("%s->%s\t", val, newval)
		}
		writer.Write(newRecord)
		//fmt.Printf("\n")
	}
	//fmt.Printf("\n%d records\n", len(rawCSVdata[1:]))
	//fmt.Println()
	//fmt.Println()
	writer.Flush()
	f, _ := os.Create("dat2.csv")
	defer f.Close()
	b.WriteTo(f)

}
