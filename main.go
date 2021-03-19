package main

import (
	"fmt"
)

func main() {
	fmt.Println("Chas Yo")
	x := Multp()
	fmt.Println("x= ", x)
	s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort1d.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort2d.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort_1line.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort_empty.csv", 1)
	s = s
	fmt.Println(s)

}
