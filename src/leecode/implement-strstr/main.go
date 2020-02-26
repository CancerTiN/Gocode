package main

import (
	"fmt"
	"strings"
)

func getPartialMatchValue(s string) (v int) {
	m := make(map[string]bool)
	for i := 1; i < len(s); i++ {
		for _, str := range []string{s[:i], s[i:]} {
			_, ok := m[str]
			if ok {
				m[str] = true
			} else {
				m[str] = false
			}
		}
	}
	for str, b := range m {
		if b && len(str) > v {
			v = len(str)
		}
	}
	return
}

func partialMatchValue(s string) (v int) {
	for i := 1; i < len(s); i++ {
		if s[:i] == s[len(s)-i:] {
			v = len(s[:i])
		}
	}
	return
}

func calNext(s string) (slice []int) {
	for i := 1; i <= len(s); i++ {
		slice = append(slice, partialMatchValue(s[:i]))
	}
	return
}

func kmp(s, p string) (idx int) {
	if p == "" {
		return
	}
	idx = -1
	if len(s) == len(p) {
		if s == p {
			return 0
		} else {
			return
		}
	}
	next := calNext(p)
	for i := 0; i <= len(s)-len(p); {
		k, j := i, 0
		for k < len(s) && j < len(p) {
			if s[k] == p[j] {
				k++
				j++
			} else {
				break
			}
		}
		if j == len(p) {
			idx = i
			break
		}
		if j == 0 {
			i++
		} else {
			i += j - next[j-1]
		}
	}
	return
}

func strStr(haystack string, needle string) int {
	return kmp(haystack, needle)
}

func main() {
	fmt.Println(partialMatchValue("ABCDAB"))
	_ = "01234567890123456789012"
	s := "BBC ABCDAB ABCDABCDABDE"
	q := "ABCDABD"
	fmt.Println(calNext(q))
	fmt.Println(kmp(s, q))
	fmt.Println(strings.Index(s, q))
	fmt.Println(strStr(s, q))
}
