package main

import (
	"euler"
	"math/big"
	"fmt"
)

func p002() {
	fibs := euler.FibonacciGenerator()
	f := euler.FilterDivisor(fibs, 2)
	
	limit := big.NewInt(4e6)
	number := big.NewInt(0)
	sum := big.NewInt(0)
	
	for {
		number.Set(<-f)
		if number.Cmp(limit) == 1 { break }
		sum.Add(sum, number)
	}
	
	fmt.Println(sum)
}