package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{-1, 8, -3, 8, -1}
	fmt.Println(findMin(0, len(arr)-1, arr))
}
func cross(changes []int, first, mid, last int) (int, int, int) {
	var leftsum, sum, maxleft, rightsum, maxright int
	for i := mid; i >= first; i-- {
		sum += changes[i]
		if sum < leftsum {
			//for maximum change it to sum > leftsum
			leftsum = sum
			maxleft = i
		}
	}
	sum = 0
	for j := mid + 1; j <= last; j++ {
		sum += changes[j]
		if sum < rightsum {
			//for maximum change it to sum > rightsum
			rightsum = sum
			maxright = j
		}
	}
	return maxleft, maxright, leftsum + rightsum
}
func findMin(first, last int, changes []int) (int, int, int) {
	if first == last {
		return first, last, changes[0]
	} else {
		mid := (last + first) / 2
		leftlow, lefthigh, leftsum := findMin(first, mid, changes)
		rightlow, righthigh, rightsum := findMin(mid+1, last, changes)
		crosslow, crosshigh, crosssum := cross(changes, first, mid, last)
		if leftsum < rightsum && leftsum < crosssum {
			//for maximum change it to leftsum >= rightsum && leftsum >= crosssum
			return leftlow, lefthigh, leftsum
		} else if rightsum < leftsum && rightsum < crosssum {
			//for maximum change it to rightsum >= leftsum && rightsum >= crosssum
			return rightlow, righthigh, rightsum
		} else {
			return crosslow, crosshigh, crosssum
		}
	}
}
