package main

import "fmt"

func binaryFind(x, l, r int) (t int) {
	m := (l + r + 1) >> 1
	if x < m*m {
		t = binaryFind(x, l, m-1)
	} else if (m+1)*(m+1) < x {
		t = binaryFind(x, m, r)
	} else {
		t = m
		if (m+1)*(m+1) == x {
			t++
		}
	}
	return
}

func mySqrt(x int) int {
	return binaryFind(x, 0, x>>1+1)
}

func main() {
	for _, x := range []int{4, 8, 9, 8192, 6, 2147483647} {
		fmt.Println(x, mySqrt(x))
	}
}

func otherSqrt(x int) int {
	left := 0
	right := x>>1 + 1
	for left < right {
		middle := (left + right + 1) >> 1
		if x < middle*middle {
			right = middle - 1
		} else {
			left = middle
		}
	}
	return left
}
