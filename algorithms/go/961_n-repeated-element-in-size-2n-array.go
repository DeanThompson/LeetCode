package main

import "fmt"

// 961. N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
//
// In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.
//
// Return the element repeated N times.

func main() {
	for _, nums := range [][]int{
		{1, 2, 3, 3},
		{2, 1, 2, 5, 3, 2},
		{5, 1, 5, 2, 5, 3, 5, 4},
	} {
		rv := repeatedNTimes(nums)
		fmt.Println(nums, rv)
	}
}

// 有 N 个唯一元素和 1 个元素重复 N+1 次，所以只要出现 2 次就是目标
func repeatedNTimes(A []int) int {
	table := make(map[int]struct{})
	for _, n := range A {
		if _, ok := table[n]; ok {
			return n
		}
		table[n] = struct{}{}
	}
	return 0
}
