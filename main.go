package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(0)
	zero := big.NewInt(0)

	// Calculate 100 factorial

	I := big.NewInt(1)
	var i int64
	var j int64
	var rj float64
	var fact float64

	start := time.Now()
	for i = 0; i < 700; i++ {
		a = big.NewInt(1)
		for j = 1; j <= 170; j++ {
			I = big.NewInt(j)
			a = a.Mul(a, I)
		}
	}
	fmt.Println("Factorial      = ", a)
	fmt.Println("Factorial Time = ", time.Since(start))

	start = time.Now()
	for i = 0; i < 700000; i++ {
		rj = 1.0
		fact = 1.0
		for j = 1; j <= 170; j++ {
			rj = float64(j)
			fact = fact * rj
		}
	}
	fmt.Println("Factorial      = ", fact)
	fmt.Println("Factorial Time = ", time.Since(start))
	b = b.Add(a, zero) // Copy bigint a into bigint b

}
