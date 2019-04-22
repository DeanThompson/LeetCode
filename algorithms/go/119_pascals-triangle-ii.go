package main

import (
	"fmt"
)

// https://leetcode.com/problems/pascals-triangle-ii/
// 119. Pascal's Triangle II
// Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
//
// Note that the row index starts from 0.

func main() {
	fmt.Println(getRow(33))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	prevRow := make([]int, rowIndex)
	result := make([]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		result[0] = 1
		for j := 1; j < i; j++ {
			result[j] = prevRow[j-1] + prevRow[j]
		}
		result[i] = 1
		copy(prevRow, result[:i+1])
	}
	return result
}
