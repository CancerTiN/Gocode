package main

import "fmt"

func findMaxCrossingSubarray(slice []int, low, mid, high int) (maxLeft, maxRight, crossSum int) {
	var sum int
	maxLeft = mid
	leftSum := slice[maxLeft]
	sum = 0
	for i := mid; i >= low; i-- {
		sum += slice[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	maxRight = mid + 1
	rightSum := slice[maxRight]
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += slice[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	crossSum = leftSum + rightSum
	return
}

func findMaximumSubarray(slice []int, low, high int) (int, int, int) {
	mid := (low + high) / 2
	if low == high {
		return low, high, slice[mid]
	}
	leftLow, leftHigh, leftSum := findMaximumSubarray(slice, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(slice, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(slice, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if leftSum <= rightSum && crossSum <= rightSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func main() {
	slice := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	low, high, maxSum := findMaximumSubarray(slice, 0, len(slice)-1)
	fmt.Printf("low: %#v, high: %#v, maximum sum: %#v\n", low, high, maxSum)
	fmt.Println("maximum subarray:", slice[low:high+1])
}
