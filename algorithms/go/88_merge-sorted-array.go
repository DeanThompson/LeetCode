package main

import "fmt"

// https://leetcode.com/problems/merge-sorted-array/
// 88. Merge Sorted Array
// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
// Note:
//
//   - The number of elements initialized in nums1 and nums2 are m and n respectively.
//   -You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.

func main() {
	nums1 := []int{1, 0, 0, 0, 0, 0}
	m := 1
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])
	fmt.Println(nums1, nums1Copy)

	i, j, idx := 0, 0, 0
	for i < m && j < n {
		if nums1Copy[i] <= nums2[j] {
			nums1[idx] = nums1Copy[i]
			i++
		} else {
			nums1[idx] = nums2[j]
			j++
		}
		idx++
	}

	for i < m {
		nums1[idx] = nums1Copy[i]
		i++
		idx++
	}
	for j < n {
		nums1[idx] = nums2[j]
		j++
		idx++
	}
}
