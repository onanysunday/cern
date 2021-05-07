package main

import (
	"fmt"
	"math/rand"
)

func GenerateRandomMatrix(height int, width int, PctFull float32) [][]float32 {
	fmt.Printf("Gen Random:   height = %v   width = %v\n", height, width)

	mat := make([][]float32, 0)

	for row := 0; row < height; row++ {
		lin := make([]float32, width)
		mat = append(mat, lin)
		//fmt.Printf("&mat[%v][0] = %v\n", row, &mat[row][0])
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if rand.Float32() < PctFull {
				mat[row][col] = rand.Float32()
			}
		}
	}

	var k int = 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mat[i][j] > 0 {
				k++
			}
		}
	}
	fmt.Printf("k = %v\n", k)

	//fmt.Printf("&mat[63][31] = %v\n", &mat[63][31])
	//fmt.Printf("&mat[0][15] = %v\n", &mat[0][15])
	//fmt.Printf("&mat[0][31] = %v\n", &mat[0][31])
	//fmt.Printf("&mat[1][0] = %v\n", &mat[1][0])
	//fmt.Printf("&mat[16][0] = %v\n", &mat[16][0])
	//fmt.Printf("&mat[63][0] = %v\n", &mat[63][0])
	return mat
}
