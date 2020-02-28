package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := 0, 0
	nums := make([]int, m+n)
	for i, _ := range nums {
		if i1 < m && i2 < n {
			if nums1[i1] < nums2[i2] {
				nums[i] = nums1[i1]
				i1++
			} else {
				nums[i] = nums2[i2]
				i2++
			}
		} else if i1 < m {
			nums[i] = nums1[i1]
			i1++
		} else if i2 < n {
			nums[i] = nums2[i2]
			i2++
		}
	}
	for i, _ := range nums {
		nums1[i] = nums[i]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Println(nums1)
	fmt.Println(nums2)
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
