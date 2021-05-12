package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func WriteCsvFile_float32(mat [][]float32, filepath string, comma rune, wformat byte) error {
	//  fmt can be:  e, E, f, g, G, x, X   (see FormatFloat in golang documentation)

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
			s := strconv.FormatFloat(v, wformat, -1, 32)
			str += s + ","
		}
		str = strings.TrimSuffix(str, ",")
		str += "\n"
		//fmt.Println("str = ", str)
		lth, err := f.WriteString(str)
		_ = lth
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil

}
