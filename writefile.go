package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"os"
	//"strconv"
)

func WriteCsvFile(mat [][]float32, filepath string, comma rune) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()

	fmt.Println("Done Writing File")
	return nil

}
