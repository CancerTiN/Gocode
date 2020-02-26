package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return i
}

func main() {
	for _, nums := range [][]int{{3, 2, 2, 3}, {0, 1, 2, 2, 3, 0, 4, 2}} {
		fmt.Println(nums)
		fmt.Println(removeElement(nums, 2))
		fmt.Println(nums)
	}
}
