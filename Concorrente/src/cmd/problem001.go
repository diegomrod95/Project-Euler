package main

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we 
// get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000. 

import (
	"fmt"
	"euler"
)

func p001() {
	c1 := euler.MultipleGenerator(3, 5)
	var soma int64 = 0
	
	for {
		select {
		case num := <-c1:
			if num >= 1000 {
				fmt.Println(soma)
				return
			}
			soma += num
		}
	}
}

