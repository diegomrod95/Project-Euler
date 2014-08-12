package euler

import (
	"math/big"
)

// Gera múltiplos de um determinado número e os envia pelo chanel retornado.
func MultipleGenerator( nums ...int64 ) <-chan int64 {
	out := make(chan int64)
	go func() {
		var i int64 = 1
		for {
			is := false
			for _, val := range nums {
				if i % val == 0 {
					is = true
				}
			}
			if is { out <- i }
			i += 1
		}
		close(out)
	}()
	return out
}

// Gera os fatores de um numero
func FactorsGenerator( number *big.Int ) <-chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		n := big.NewInt(1)
		zero := big.NewInt(0)
		one := big.NewInt(1)
		for {
			q := big.NewInt(1).Quo(n, number)
			if q.Cmp(zero) == 0 {
				out <- n
			}
			n.Add(n, one)
		}
		close(out)
	}()
	return out
}

// Gera números fibonacci
func FibonacciGenerator() <-chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		a := big.NewInt(0)
		b := big.NewInt(1)
		for {
			c := big.NewInt(0)
			c.Set(a)
			a.Set(b)
			b.Add(b, c)
			out <- c
		}
		close(out)
	}()
	return out
}