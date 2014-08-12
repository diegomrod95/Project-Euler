package main 

import (
	"euler"
	"math"
	"math/big"
	"fmt"
)

func p003() {
	gen := euler.FactorsGenerator(big.NewInt(int64(math.Sqrt(600851475143))))
	prm := euler.FilterPrimes(gen)
	
	for i := 0; i < 1000; i += 1 {
		fmt.Println(<-prm)
	}
}

