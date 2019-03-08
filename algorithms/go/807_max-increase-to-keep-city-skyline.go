package main

import "fmt"

// 807. Max Increase to Keep City Skyline
// https://leetcode.com/problems/max-increase-to-keep-city-skyline/

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	rv := maxIncreaseKeepingSkyline(grid)
	fmt.Println(rv)
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	length := len(grid[0])
	topBottom := make([]int, length)
	leftRight := make([]int, length)

	for i := 0; i < length; i++ {
		lrMax := grid[i][0]
		tbMax := grid[0][i]
		for j := 0; j < length; j++ {
			if grid[i][j] > lrMax {
				lrMax = grid[i][j]
			}
			if grid[j][i] > tbMax {
				tbMax = grid[j][i]
			}
		}
		leftRight[i] = lrMax
		topBottom[i] = tbMax
		fmt.Println(grid[i], leftRight, topBottom)
	}

	rv := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			grid[i][j] = min(leftRight[i], topBottom[j])
			rv += min(leftRight[i], topBottom[j]) - grid[i][j]
		}
		fmt.Println(grid[i])
	}

	return rv
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
