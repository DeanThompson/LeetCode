package main

import "fmt"

// 461. Hamming Distance
// https://leetcode.com/problems/hamming-distance/
//
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
// Given two integers x and y, calculate the Hamming distance.

func main() {
	x, y := 1, 4
	dist := hammingDistance(x, y)
	fmt.Println(x, y, dist)
}

// 算法来自维基百科: https://en.wikipedia.org/wiki/Hamming_distance
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
