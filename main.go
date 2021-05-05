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
	//filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx.txt"
	//filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx_SCALED.txt"
	filename := "C:/Users/onany/OneDrive/Documents/Data/BigFish/Y3Y_NEED_Idx_100.txt"
	skiplines := 1
	//	records := readCsvFile2d_64(filename, ',', skiplines)
	records := readCsvFile2d_float32(filename, ',', skiplines)
	recordSize := len(records)
	fmt.Println("#Records = ", recordSize)
	fmt.Printf("records type ' %T\n", records)
	fmt.Println(records[0])

	rmat := GenerateRandomMatrix(64, 32, 0.5)
	_ = rmat
	Writepath := "C:/Users/onany/OneDrive/Documents/krud/matout.txt"
	err := WriteCsvFile(rmat, Writepath, ',')
	_ = err

	PrintMemUsage()

}
