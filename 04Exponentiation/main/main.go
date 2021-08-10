package main

import (
	"fmt"
)

func main() {
	var num, pow int
	fmt.Scan(&num, &pow)
	ans := Power(num, pow)
	fmt.Println(ans)
}

//We've done this because we wanted an algorithm with lower growth order than n
//if we simply used the naive algorithm in the normal way it would be O(n)(big o of n)
func Power(num int, pow int) int {
	// The idea is to divide the problem into two subproblems and then solve it piece by piece
	//and the combine it back together.
	if pow == 0 {
		// here is where our split job is done
		return 1
	} else if pow%2 == 0 {
		if pow < 2 {
			return num
		}
		split := Power(num, pow/2)
		return split * split
	} else {
		if pow < 2 {
			return num
		}
		split := Power(num, (pow-1)/2)
		return split * split * num
	}
}
