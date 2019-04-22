package main

import "fmt"

// https://leetcode.com/problems/pascals-triangle/
// 118. Pascal's Triangle
// Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.

func main() {
	printTriangle(generate(1))
	printTriangle(generate(2))
	printTriangle(generate(5))
}

func printTriangle(data [][]int) {
	for _, row := range data {
		fmt.Println(row)
	}
	fmt.Println()
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		row[i] = 1
		result[i] = row

	}
	return result
}
