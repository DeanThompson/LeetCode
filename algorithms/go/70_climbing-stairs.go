package main

import "fmt"

// 70. Climbing Stairs
// You are climbing a stair case. It takes n steps to reach to the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Note: Given n will be a positive integer.

func main() {
	for _, ns := range [][]int{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	} {
		fmt.Printf("n: %v, expected: %v, got: %v\n", ns[0], ns[1], climbStairs(ns[0]))
	}
}

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i < n+1; i++ {
		a, b = b, a+b
	}
	return b
}
