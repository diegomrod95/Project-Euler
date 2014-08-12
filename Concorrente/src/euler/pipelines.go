package euler

import (
	"math/big"
)

// Junção que filtra os números dvisores de num vindos de um canal e envia para
// outro canal.
func FilterDivisor( seq <-chan *big.Int, num int64 ) <-chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		m := big.NewInt(num)
		zero := big.NewInt(0)
		for {
			select {
			case n := <-seq:
				mod := big.NewInt(0)
				mod.Mod(n, m)
				if mod.Cmp(zero) == 0 {
					out <- n
				}
			}
		}
		close(out)
	}()
	return out
}

func FilterPrimes( seq <-chan *big.Int ) <-chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		for {
			select {
			case n := <-seq:
				if IsPrime(n) {
					out <- n
				}
			}
		}
		close(out)
	}()
	return out
}