package main

import (
	"fmt"
	"runtime"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	fmt.Println("Chas Yo")
	//x := Multp()
	//fmt.Println("x= ", x)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort2d.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort_1line.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort_empty.csv", 1)
	//filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx.txt"
	//filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx_100.txt"
	filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx_SCALED.txt"
	skiplines := 1
	//	records := readCsvFile2d_64(filename, ',', skiplines)
	records := readCsvFile2d_float32(filename, ',', skiplines)
	recordSize := len(records)
	fmt.Println("#Records = ", recordSize)
	fmt.Printf("records type ' %T\n", records)
	fmt.Println(records[0])

	PrintMemUsage()

	var sum float32
	for row := 0; row < 5594672; row++ {
		sum += records[row][5]
	}
	fmt.Println("Sales Sum = ", sum)

}
