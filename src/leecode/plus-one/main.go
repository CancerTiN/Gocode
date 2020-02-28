package main

import "fmt"

func plusOne(digits []int) []int {
	carry := false
	stretch := false
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i]++
		} else if carry {
			digits[i]++
		}
		if digits[i] == 10 {
			carry = true
			digits[i] = 0
		} else {
			carry = false
		}
		if i == 0 && carry {
			stretch = true
		}
	}
	if stretch {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	for _, digits := range [][]int{{1, 2, 3}, {4, 3, 2, 1}, {9}, {9, 9, 9}} {
		fmt.Println(digits)
		fmt.Println(plusOne(digits))
	}
}
