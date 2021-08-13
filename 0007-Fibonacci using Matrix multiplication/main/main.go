package main

import "fmt"

func main() {

	var n int = 6
	a := [][]int{
		{1, 1},
		{1, 0},
	}
	var nth [][]int = a
	//The whole procedure is just like the simple matrix multiplication
	//We only use this for loop to calculate the Fibonacci number that we need
	//and by the way the matrix that we get is [Fn+1 , Fn]
	//										   [Fn , Fn-1]
	for i := 0; i < n; i++ {
		nth = matrixMult(nth, a)
	}
	// fmt.Println(s)

	//let's say what if our first Fibo numbers were not 1 and 1
	//so what will be our nth number? Here we only need to multiply the
	//result by our first numbers.
	//a1 and a2 are our first numbers and we can use whatever we want and the result is [an+1]
	//																					[ an ]

	a0 := 5
	a1 := 6
	first := [][]int{
		{a1},
		{a0},
	}
	nthz := matrixMult(nth, first)
	fmt.Println(nthz)
}

func matrixMult(a [][]int, b [][]int) (c [][]int) {
	c = make([][]int, len(a))
	for i := range c {
		c[i] = make([]int, len(b[0]))
	}
	if len(a[0]) != len(b) {
		fmt.Println("can't multiply these two matrices")
		return c
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				c[i][j] = c[i][j] + a[i][k]*b[k][j]
			}
		}
	}
	return c
}
