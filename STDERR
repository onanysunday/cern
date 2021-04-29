package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filepath string, comma rune, skiplines int) ([][]string, int) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()
	// 1032
	csvReader := csv.NewReader(f)
	csvReader.Comma = comma       // can be ',' or '\t' or '|' or whatever
	csvReader.FieldsPerRecord = 0 // tell csvReader to return the number of fields separated by csvReader.Comma
	//csvReader.FieldsPerRecord = 4
	for i := 0; i < skiplines; i++ {
		topline, err1 := csvReader.Read()
		_ = topline
		if err1 != nil {
			log.Fatal("Could not read top line "+filepath, err1)
		}
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filepath, err)
	}
	//fmt.Println("Size of array = ", len(records))
	//fmt.Println("Size of arrayz = ", len(records[0]))

	return records, csvReader.FieldsPerRecord
}

func readCsvFile2d_int(filepath string, comma rune, skiplines int) [][]int {
	records, fieldsPerRecord := readCsvFile(filepath, comma, skiplines)
	recordSize := len(records)
	//fmt.Println("Records", records)
	fmt.Println("fieldsPerRecord", fieldsPerRecord)
	fmt.Println("recordSize", recordSize)

	mat := make([][]int, 0)
	var err1 error
out:
	for row := 0; row < recordSize; row++ {
		lin := make([]int, fieldsPerRecord)
		for col := 0; col < fieldsPerRecord; col++ {
			lin[col], err1 = strconv.Atoi(records[row][col])
			if err1 != nil {
				fmt.Println("Error reading line #", row, err1)
				break out
			}
		}
		mat = append(mat, lin)
		//fmt.Println(lin)
	}

	fmt.Println("mat[0][0] = ", &mat[0][0])
	fmt.Println("mat[1][0] = ", &mat[1][0])
	fmt.Printf("mat Type = %T ", mat)
	return mat
}

func readCsvFile2d_64(filepath string, comma rune, skiplines int) [][]int64 {
	records, fieldsPerRecord := readCsvFile(filepath, comma, skiplines)
	recordSize := len(records)
	//fmt.Println("Records", records)
	fmt.Println("fieldsPerRecord", fieldsPerRecord)
	fmt.Println("recordSize", recordSize)

	mat := make([][]int64, 0)
	var err1 error
out:
	for row := 0; row < recordSize; row++ {
		lin := make([]int64, fieldsPerRecord)
		for col := 0; col < fieldsPerRecord; col++ {
			lin[col], err1 = strconv.ParseInt(records[row][col], 10, 64)
			if err1 != nil {
				fmt.Println("Error reading line #", row, err1)
				break out
			}
		}
		mat = append(mat, lin)
		//fmt.Println(lin)
	}

	fmt.Println("&mat[0][0] = ", &mat[0][5])
	fmt.Println("mat[0][0] = ", mat[0][5])
	//fmt.Println("mat[1][0] = ", &mat[1][0])
	fmt.Printf("mat Type = %T\n", mat)
	return mat
}

/*
func zzz_readCsvFile1d(filepath string, skiplines int) []string {
	records, fieldsPerRecord := readCsvFile(filepath, skiplines)

	recordSize := len(records)
	fieldsPerRecord = fieldsPerRecord

	var Rec1D = make([]string, recordSize)
	for i := 0; i < len(Rec1D); i++ {
		Rec1D[i] = records[i][0]
	}
	//fmt.Println("rec = ", Rec1D)
	return Rec1D
}

func zzz_readCsvFile1d_int(filepath string, skiplines int) []int {
	records, fieldsPerRecord := readCsvFile(filepath, skiplines)

	recordSize := len(records)
	fieldsPerRecord = fieldsPerRecord

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
*/
