package main

import (
	"fmt"
)

func main() {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 1, 1, 1, 1}
	result := true
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			check := false
			for j := i; j < len(a); j++ {
				if a[i] == b[j] {
					b[i], b[j] = b[j], b[i]
					break
				} else if j == len(a)-1 {
					result = false
					check = true
					break
				}
			}
			if check {
				break
			}
		}
	} else {
		result = false
	}
	fmt.Println(result)
}
