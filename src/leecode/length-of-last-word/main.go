package main

import "fmt"

func lengthOfLastWord(s string) int {
	isSuffixSpace := true
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if isSuffixSpace {
				continue
			} else {
				break
			}
		} else {
			isSuffixSpace = false
			n++
		}
	}
	return n
}

func main() {
	for _, s := range []string{"Hello World", "a "} {
		fmt.Println(s, lengthOfLastWord(s))
	}
}
