package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, s := range []string{"11", "1", "1010", "1011"} {
		i64, err := strconv.ParseUint(s, 2, 0)
		fmt.Println(s, i64, err)
	}

	fmt.Println(2, int('3'-'0'), '8'-'0')
	for _, b := range []byte{0, 1, 2, 3} {
		fmt.Println(b, b%2, b%2+'0', string(b%2+'0'))
	}
}
