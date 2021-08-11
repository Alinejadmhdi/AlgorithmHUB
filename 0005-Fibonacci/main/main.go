package main

import "fmt"

var calcIt, i int
var first, sec bool
var cals []int

//Here is the famous recursive fibonacci
// func fibonacci(n int) int {
// 	if n == 0 {
// 		return 0
// 	} else if n == 1 {
// 		return 1
// 	}
// 	return fibonacci(n-1) + fibonacci(n-2)
// }

func main() {
	fmt.Println(fibonacci(7))
}

// ok, when I looked at the normal recursive Fibonacci first I said
// why we should call the function twice when we already found the f(n-2)
// while we called f(n-1), so here we are, I did it with only one call of the function
func fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	calcIt = fibonacci(n - 1)
	// the main part of my code happens below where you can see the first time and  the second time
	//my slice was not long enough to call while I wanted to use my last indexes to
	//do the summation
	cals = append(cals, calcIt)
	// for the first index
	if !first && !sec {
		i = 0
		first = true
		return cals[i]
		// for the second index
	} else if !sec {
		i = 1
		sec = true
	}
	// for moving through our slice
	i++
	return cals[i-1] + cals[i-2]
}
