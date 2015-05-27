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
	csvfile, err := os.Open(os.Args[1])
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
	m := make(map[string]func(rune) rune)
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
			}
			newRecord[i] = newval
			//fmt.Printf("%s->%s\t", val, newval)
		}
		writer.Write(newRecord)
		//fmt.Printf("\n")
	}
	writer.Flush()
	f, _ := os.Create("data.csv")
	defer f.Close()
	b.WriteTo(f)
	fmt.Println("Output written to data.csv")
}
