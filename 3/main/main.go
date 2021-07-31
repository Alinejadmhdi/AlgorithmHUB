package main

import (
	"fmt"
	"os"
)

var found bool = false

func main() {
	//In case you want to take the array from user.
	// array := takearr()
	array := []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 200}
	item := 199
	//In case you want to take the item from user.
	// var item int
	// fmt.Println("Enter Your Number!")
	// fmt.Scan(&item)
	binSearch(array, item)
}

func binSearch(arr []int, item int) {
	if len(arr) == 1 {
		return
	}
	m := len(arr) / 2
	if arr[m] == item {
		fmt.Println("Yes!")
		found = true
	}
	if arr[m] < item {
		binSearch(arr[m:], item)
	} else if arr[m] > item {
		binSearch(arr[:m], item)
	}
	if !found {
		fmt.Println("Nope")
		os.Exit(1)
	}
}

// func takearr() []int {
// 	fmt.Println("We need a sorted array!")
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
