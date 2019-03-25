package main

import "fmt"

// 922. Sort Array By Parity II
// https://leetcode.com/problems/sort-array-by-parity-ii/
//
// Given an array A of non-negative integers, half of the integers in A are odd,
// and half of the integers are even.
//
// Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.
//
// You may return any answer array that satisfies this condition.

func main() {
	A := []int{4, 2, 5, 7}
	fmt.Println([]int{4, 2, 5, 7}, sortArrayByParityII(A))

	A = []int{1, 2, 3, 4, 5, 6}
	fmt.Println([]int{4, 2, 5, 7}, sortArrayByParityII(A))
}

func sortArrayByParityII(A []int) []int {
	result := make([]int, len(A))
	evenIdx, oddIdx := 0, 1
	for _, v := range A {
		if v%2 == 0 {
			result[evenIdx] = v
			evenIdx += 2
		} else {
			result[oddIdx] = v
			oddIdx += 2
		}
	}
	return result
}
