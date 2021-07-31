package main

import (
	"fmt"
)

func main() {
	// In case you want to take the array from user.
	// arr := takearr()
	arr := []int{1, 5, 6, 8, 2, 3, 1}
	aarr := mergeSort(arr)
	fmt.Println(aarr)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	m1 := arr[:len(arr)/2]
	m2 := arr[len(arr)/2:]
	l := mergeSort(m1)
	r := mergeSort(m2)
	return merge(l, r)
}

func merge(l []int, r []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			final = append(final, l[i])
			i++
		} else {
			final = append(final, r[j])
			j++
		}
	}
	for ; i < len(l); i++ {
		final = append(final, l[i])
	}
	for ; j < len(r); j++ {
		final = append(final, r[j])
	}
	return final
}

// func takearr() []int {
// 	var n int
// 	fmt.Scan(&n)
// 	arr := make([]int, n)
// 	if n <= 1 {
// 		fmt.Println("Done!", arr)
// 		os.Exit(1)
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i])
// 	}
// 	return arr
// }
