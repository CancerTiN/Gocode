package main

import "fmt"

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for ; j >= 0; j-- {
			if slice[j] > key {
				slice[j+1] = slice[j]
			} else {
				break
			}
		}
		slice[j+1] = key
	}
}

func main() {
	slice := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(slice)
	insertionSort(slice)
	fmt.Println(slice)
}
