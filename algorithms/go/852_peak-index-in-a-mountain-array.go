package main

import "fmt"

// 852. Peak Index in a Mountain Array
// https://leetcode.com/problems/peak-index-in-a-mountain-array/

func main() {
	A := []int{0, 1, 0}
	fmt.Println(A, peakIndexInMountainArray(A))

	A = []int{0, 2, 1, 0}
	fmt.Println(A, peakIndexInMountainArray(A))

	A = []int{1, 2, 3, 4, 5, 1}
	fmt.Println(A, peakIndexInMountainArray(A))
}

func peakIndexInMountainArray(A []int) int {
	return binarySearch(A)
}

func linearScan(A []int) int {
	rv := 0
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			rv = i
			break
		}
	}
	return rv
}

func binarySearch(A []int) int {
	if len(A) <= 2 {
		return len(A) - 1
	}
	i := len(A) / 2
	if (A[i-1] < A[i]) && (A[i] > A[i+1]) {
		return i
	}
	if A[i-1] > A[i] {
		return binarySearch(A[:i])
	}
	return i + binarySearch(A[i:])
}
