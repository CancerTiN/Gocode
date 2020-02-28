package main

import "fmt"

func main() {
	for _, n := range []int{2, 4, 5, 8, 9, 55, 100} {
		fmt.Println(n, n<<1, n>>1, n<<1>>2, n>>1<<2)
	}
}
