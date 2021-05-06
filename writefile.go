package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"os"
	//"strconv"
)

func WriteCsvFile(mat [][]float32, filepath string, comma rune) error {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()

	lth, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println("File Length = ", lth)

	fmt.Println("Done Writing File")
	return nil

}
