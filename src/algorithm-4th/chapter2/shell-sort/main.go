package main

import "fmt"

func shellSort(slice []byte) {
	length, h := len(slice), 1
	for h < length/3 {
		h = 3*h - 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && slice[j] < slice[j-h]; j -= h {
				slice[j], slice[j-h] = slice[j-h], slice[j]
			}
		}
		h = h / 3
	}
}

func main() {
	slice := []byte{'S', 'H', 'E', 'L', 'L', 'S', 'O', 'R', 'T', 'E', 'X', 'A', 'M', 'P', 'L', 'E'}
	fmt.Println(string(slice))
	shellSort(slice)
	fmt.Println(string(slice))
}
