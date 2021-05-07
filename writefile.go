package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func WriteCsvFile(mat [][]float32, filepath string, comma rune) error {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Unable to Create Text File "+filepath, err)
	}
	defer f.Close()

	nrows := len(mat)
	fmt.Println("nrows = ", nrows)
	for row := 0; row < nrows; row++ {
		record := mat[row]
		_ = record
		lastcol := len(mat[row]) - 1
		var str string = ""
		for col := 0; col <= lastcol; col++ {
			v := float64(mat[row][col])
			s := strconv.FormatFloat(v, 'g', -1, 32)
			str += s
			if col != lastcol {
				str += ","
			}
		}
		fmt.Println("str = ", str)
		//fmt.Println("record = ", record)
	}

	lth, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("File Length = ", lth)

	fmt.Println("Done Writing File")
	return nil

}
