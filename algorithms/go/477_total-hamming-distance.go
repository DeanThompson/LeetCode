package main

import "fmt"

// 477. Total Hamming Distance
// https://leetcode.com/problems/total-hamming-distance/
//
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
// Now your job is to find the total Hamming distance between all pairs of the given numbers.

func main() {
	nums := []int{4, 14, 2}
	rv := totalHammingDistance(nums)
	fmt.Println(nums, rv)
}

func totalHammingDistance(nums []int) int {
	return iterate(nums)
}

func iterate(nums []int) int {
	rv := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			rv += hammingDistance(nums[i], nums[j])
		}
	}
	return rv
}

func hammingDistance(x int, y int) int {
	dist := 0
	val := x ^ y

	// Count the number of bits set
	for val != 0 {
		dist++
		val &= val - 1
	}
	return dist
}

// TODO: 使用 bitmap 优化
// 参考 https://leetcode.com/problems/total-hamming-distance/discuss/244831/Bits-trick-very-simple-solution-approach.
