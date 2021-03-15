package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsvFile(filepath string, skiplines int) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for i := 0; i < skiplines; i++ {
		topline, err1 := csvReader.Read()
		topline = topline
		if err1 != nil {
			log.Fatal("Could not read top line "+filepath, err1)
		}
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filepath, err)
	}
	//k :=  (records[0][0])
	//fmt.Println("k = ", k)

	return records
}
