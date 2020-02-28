package main

import "fmt"

var m = map[int]int{1: 1, 2: 2}

func climbStairs(n int) int {
	if n < 3 {
		return m[n]
	} else {
		r, ok := m[n]
		if !ok {
			r = climbStairs(n-1) + climbStairs(n-2)
			m[n] = r
		}
		return r
	}
}

func main() {
	for _, n := range []int{1, 2, 3, 4, 44} {
		fmt.Println(n, climbStairs(n))
	}
}
