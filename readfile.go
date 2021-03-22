package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filepath string, skiplines int) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()
	// 1032
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
	//fmt.Println("Size of array = ", len(records))
	//fmt.Println("Size of arrayz = ", len(records[0]))

	return records
}

func readCsvFile2d(filepath string, skiplines int) [][]string {
	records := readCsvFile(filepath, skiplines)

	return records
}

func readCsvFile1d(filepath string, skiplines int) []string {
	records := readCsvFile(filepath, skiplines)

	recordSize := len(records)

	var Rec1D = make([]string, recordSize)
	for i := 0; i < len(Rec1D); i++ {
		Rec1D[i] = records[i][0]
	}
	//fmt.Println("rec = ", Rec1D)
	return Rec1D
}

func readCsvFile1d_int(filepath string, skiplines int) []int {
	records := readCsvFile(filepath, skiplines)

	recordSize := len(records)

	var Rec1D = make([]int, recordSize)
	for i := 0; i < len(Rec1D); i++ {
		//Rec1D[i] = strconv.Atoi(records[i][0])
		x, serr := strconv.Atoi(records[i][0])
		if serr != nil {
			fmt.Println("Convert Error", records[i][0])
		}

		serr = serr
		Rec1D[i] = x
	}
	fmt.Println("rec = ", Rec1D)
	return Rec1D
}
