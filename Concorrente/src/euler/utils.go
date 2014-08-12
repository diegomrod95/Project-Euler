package euler

import (
	"math/big"
)

func IsPrime( number *big.Int ) bool {
	if big.NewInt(0).Mod(number, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return false
	}
	if big.NewInt(0).Mod(number, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
		return false
	}
	if big.NewInt(0).Mod(number, big.NewInt(5)).Cmp(big.NewInt(2)) == 0 {
		return false
	}
	return true
}
