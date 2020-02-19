package main

import "fmt"

func isUnique(astr string) bool {
	m := make(map[rune]struct{})
	for _, b := range astr {
		_, ok := m[b]
		if ok {
			return false
		} else {
			m[b] = struct{}{}
		}
	}
	return true
}

func main() {
	for _, s := range []string{"leetcode", "abc"} {
		fmt.Println(s, isUnique(s))
	}
}
