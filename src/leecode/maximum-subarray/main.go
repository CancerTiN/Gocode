package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum <= 0 && curSum < nums[i] {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	slice := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(slice, maxSubArray(slice))
}
