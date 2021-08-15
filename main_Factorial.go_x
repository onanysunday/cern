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

	start := time.Now()
	for i = 0; i < 700000; i++ {
		a = big.NewInt(1)
		for j = 1; j <= 100; j++ {
			I = big.NewInt(j)
			a = a.Mul(a, I)
		}
	}
	fmt.Println("Factorial Time = ", time.Since(start))

	b = b.Add(a, zero) // Copy bigint a into bigint b

	fmt.Printf("Vals:  %v %v\n", a, b)   // a = *big.Int, b = big.Int
	fmt.Printf("Types: %T %T\n", a, b)   // a = *big.Int, b = big.Int
	fmt.Printf("Addr:  %v %v\n", &a, &b) // a = *big.Int, b = big.Int

	// Unwind the factorial
	start = time.Now()
	for i = 0; i < 700000; i++ {
		b = b.Add(a, zero)
		for j = 1; j <= 100; j++ {
			I = big.NewInt(j)
			b = b.Div(b, I)
		}
	}
	fmt.Println("Unfactorial Time = ", time.Since(start))
	fmt.Println("unf: a = ", a)
	fmt.Println("unf: b = ", b)
}
