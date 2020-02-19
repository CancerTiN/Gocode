package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
	m1 := make(map[int32]int)
	for _, b := range s1 {
		_, ok := m1[b]
		if ok {
			m1[b]++
		} else {
			m1[b] = 1
		}
	}
	m2 := make(map[int32]int)
	for _, b := range s2 {
		_, ok := m2[b]
		if ok {
			m2[b]++
		} else {
			m2[b] = 1
		}
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		} else {
			if v1 != v2 {
				return false
			}
		}
	}
	return true
}

func main() {
	for _, slice := range [][]string{{"abc", "bca"}, {"abc", "bad"}} {
		fmt.Println(slice[0], slice[1], CheckPermutation(slice[0], slice[1]))
	}
}
