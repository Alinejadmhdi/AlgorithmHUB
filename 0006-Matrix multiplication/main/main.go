package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{0, -1, 2},
		{4, 11, 2},
	}
	b := [][]int{
		{3, -1},
		{1, 2},
	}
	fmt.Println(matrixMult(a, b))
}

//First I made an empty array for the answer
// and then I just used 2 loops to move over matrices
// and another one to move on the mults.
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
