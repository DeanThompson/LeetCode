package main

import "fmt"

/*
977. Squares of a Sorted Array

https://leetcode.com/problems/squares-of-a-sorted-array/

Given an array of integers A sorted in non-decreasing order,
return an array of the squares of each number, also in sorted non-decreasing order.

Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
 */

func main() {
	nums := []int{0, 1, 2, 3}

	rv := sortedSquares(nums)
	fmt.Println(rv)
}

func sortedSquares(A []int) []int {
	size := len(A)
	result := make([]int, size)

	threshold := 0
	if A[size-1] <= 0 {
		threshold = size
	} else {
		for i, n := range A {
			if n >= 0 {
				threshold = i
				break
			}
		}
	}

	i, j := threshold-1, threshold
	idx := 0
	for i >= 0 && j < size {
		if 0-A[i] < A[j] {
			result[idx] = A[i] * A[i]
			i--
		} else {
			result[idx] = A[j] * A[j]
			j++
		}
		idx++
	}

	for i >= 0 {
		result[idx] = A[i] * A[i]
		idx++
		i--
	}
	for j < size {
		result[idx] = A[j] * A[j]
		idx++
		j++
	}
	return result
}
