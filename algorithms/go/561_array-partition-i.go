package main

import (
	"fmt"
	"sort"
)

// 561. Array Partition I
// https://leetcode.com/problems/array-partition-i/
//
// Given an array of 2n integers, your task is to group these integers into n pairs of integer,
// say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n
// as large as possible.

func main() {
	nums := []int{1, 4, 3, 2}
	rv := arrayPairSum(nums)
	fmt.Println(nums, rv)
}

// 升序排序，对奇数位置（第一个位置下标为 0）元素求和
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	rv := 0
	for i := 0; i < len(nums)-1; i += 2 {
		rv += nums[i]
	}
	return rv
}
