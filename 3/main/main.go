package main

import (
	"fmt"
	"os"
)

var firstRun, left, right bool = true, false, false
var index, lastItem int
var m int

func main() {
	//In case you want to take the array from the user.
	//hey if you want to use it this way just go to the end of code and simply uncomment
	//my lovely "takearr" function

	// array := takearr()
	array := []int{1, 3, 4, 9, 20, 31, 45, 63, 70, 100, 200}
	item := 8
	//In case you want to take the item from the user.

	// var item int
	// fmt.Println("Enter Your Number!")
	// fmt.Scan(&item)
	binSearch(array, item)
}

func binSearch(arr []int, item int) {
	//our journey will finish here
	if lastItem > 1 {
		fmt.Println("We Couldn't Find Your Item!")
		os.Exit(0)
		return
	}
	m = len(arr) / 2
	// m will never be under 0 so we need to know when we couldn't find the item
	if m == 0 {
		lastItem++
	}
	//calculating the index, I didn't want to use the famous type of finding its index so I
	//simply followed m, if you still don't know why I did these two operations simply
	//draw the array on the paper and check where our m goes you will get it easily
	if !firstRun {
		if left {
			index -= (len(arr) - m)
		} else if right {
			index += m
		}
	} else {
		index = m
	}
	//finding item, this part is the famous part of BinarySearch
	if arr[m] == item {
		fmt.Println("Yes! We Found Your Item It Is In Index:", index)
		os.Exit(0)
	}
	if arr[m] < item {
		firstRun = false
		right = true
		left = false
		binSearch(arr[m:], item)
	} else if arr[m] > item {
		firstRun = false
		left = true
		right = false
		binSearch(arr[:m], item)
	}
}

//this function is used to get your array from the input.

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
