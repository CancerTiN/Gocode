package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	} else {
		str := ""
		s := countAndSay(n - 1)
		n := 1
		l := s[:1]
		for i := 1; i < len(s); i++ {
			c := s[i : i+1]
			if l == c {
				n++
			} else {
				str += fmt.Sprintf("%d%s", n, l)
				n = 1
				l = c
			}
			if i == len(s)-1 {
				str += fmt.Sprintf("%d%s", n, l)
			}
		}
		return str
	}
}

func main() {
	for i := 1; i < 6; i++ {
		fmt.Println(i, countAndSay(i))
	}
}
