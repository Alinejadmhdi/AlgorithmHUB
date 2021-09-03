package main

import "fmt"

var c [][]int

/*
I wrote this I wrote this algorithm without having
the matrices' exact size. We know we can define our
2d slices without For ranges that I used, and it will
take less time.
It will make our code run inefficiently.
*/
func main() {
	a := [][]int{
		{0, -1, 9, 23},
		{4, 11, 7, 35},
		{12, 9, 88, 21},
		{2, 57, 8, 9},
	}
	b := [][]int{
		{3, -1, 95, 12},
		{1, 2, 65, 18},
		{59, 65, 47, 34},
		{16, 74, 3, 5},
	}
	c = makeArray(2 * len(a))
	fmt.Println(squareMatMult(a, b))
}

/*
This function is the primary function of my algorithm that calculates our answer.
So the main idea is to divide the problem into four submatrices and do the same
for the submatrix until our submatrix has only one item.
*/
func squareMatMult(a, b [][]int) [][]int {
	// The return condition of recursive function
	if len(a) == 1 || len(b) == 1 {
		z := makeArray(2)
		z[0][0] = a[0][0] * b[0][0]
		return z
	}
	n := len(a)
	/*
		Here we divided each matrix and filled them using the items of the main matrices.
	*/
	a11 := makeArray(n)
	a12 := makeArray(n)
	a21 := makeArray(n)
	a22 := makeArray(n)
	b11 := makeArray(n)
	b12 := makeArray(n)
	b21 := makeArray(n)
	b22 := makeArray(n)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			a11[i][j] = a[i][j]
			a12[i][j] = a[i][j+n/2]
			a21[i][j] = a[i+n/2][j]
			a22[i][j] = a[i+n/2][j+n/2]
			b11[i][j] = b[i][j]
			b12[i][j] = b[i][j+n/2]
			b21[i][j] = b[i+n/2][j]
			b22[i][j] = b[i+n/2][j+n/2]
		}
	}
	/*
		The central part of the algorithm happens here, where we partition the matrices.
		I recursively called squareMatMult a total of eight times.
		Each recursive call multiplies two n/2*n/2 matrices.
	*/
	c11 := matSum(squareMatMult(a11, b11), squareMatMult(a12, b21))
	c12 := matSum(squareMatMult(a11, b12), squareMatMult(a12, b22))
	c21 := matSum(squareMatMult(a21, b11), squareMatMult(a22, b21))
	c22 := matSum(squareMatMult(a21, b12), squareMatMult(a22, b22))

	c = makeArray(2 * len(a))

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			c[i][j] += c11[i][j]
			c[i][j+n/2] += c12[i][j]
			c[i+n/2][j] += c21[i][j]
			c[i+n/2][j+n/2] += c22[i][j]
		}
	}
	return c
}
func matSum(a, b [][]int) [][]int {
	n := len(a)
	p := makeArray(2 * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p[i][j] = a[i][j] + b[i][j]
		}
	}
	return p
}
func makeArray(n int) [][]int {
	c := make([][]int, n/2)
	for i := range c {
		c[i] = make([]int, n/2)
	}
	return c
}
