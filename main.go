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
	//s := readCsvFile1d_int("C:/Users/onany/OneDrive/Documents/krud/sort1d.csv", 0)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort2d.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort_1line.csv", 1)
	//s := readCsvFile2d("C:/Users/onany/OneDrive/Documents/krud/sort_empty.csv", 1)
	//fmt.Println(s)
	//filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx.txt"
	//filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx_100.txt"
	filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx_7x.txt"
	skiplines := 1
	//records := readCsvFile(filename, skiplines)
	records := readCsvFile2d(filename, ',', skiplines)
	recordSize := len(records)
	fmt.Println("#Records = ", recordSize)
	fmt.Printf("records type ' %T\n", records)
	fmt.Println(records[0])

	//var grid [262144][131072]byte
	//grid[1000][1000] = 99

	PrintMemUsage()

	//fmt.Println("m = ", grid[1000][1000])

}
