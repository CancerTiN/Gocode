package main

import "fmt"

func searchInsert(nums []int, target int) int {
	i := 0
	if len(nums) < 1 || target <= nums[i] {
		return i
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	for i+1 < len(nums) {
		if nums[i] <= target && target <= nums[i+1] {
			i++
			break
		}
		i++
	}
	return i
}

func main() {
	m := map[int][]int{2: {1, 3, 5, 6}, 5: {1, 3, 5, 6}, 7: {1, 3, 5, 6}}
	for target, nums := range m {
		fmt.Println(searchInsert(nums, target))
	}
}
