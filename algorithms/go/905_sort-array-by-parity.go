package main

import "fmt"

// 905. Sort Array By Parity
// https://leetcode.com/problems/sort-array-by-parity/
//
// Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
//
// You may return any answer array that satisfies this condition.
//
//
//
// Example 1:
//
// Input: [3,1,2,4]
// Output: [2,4,3,1]
// The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

func main() {
	for _, nums := range [][]int{
		{2, 4, 6, 8, 1},
		{3, 1, 2, 4, 5},
	} {
		fmt.Println(nums)
		rv := sortArrayByParity(nums)
		fmt.Println(rv)
		fmt.Println()
	}
}

func sortArrayByParity(A []int) []int {
	return quickSort(A)
}

// 选择排序，外层循环迭代 A，遇到奇数，内层循环找到下一个偶数，就地交换
func selectSort(A []int) []int {
	for idx, n := range A {
		if n%2 != 0 {
			for i := idx + 1; i < len(A); i++ {
				if A[i]%2 == 0 {
					A[idx], A[i] = A[i], A[idx]
					break
				}
			}
		}
	}
	return A
}

// 维护首尾两个指针，把头部的奇数和尾部的偶数就地交换位置
func quickSort(A []int) []int {
	head, tail := 0, len(A)-1
	for head < tail {
		if A[head]%2 == 1 && A[tail]%2 == 0 {
			A[head], A[tail] = A[tail], A[head]
		}

		if A[head]%2 == 0 {
			head++
		}
		if A[tail]%2 == 1 {
			tail--
		}
	}
	return A
}

// 两次迭代，第一次选出偶数，第二次选出奇数
func twoPass(A []int) []int {
	rv := make([]int, len(A))
	idx := 0
	for _, n := range A {
		if n%2 == 0 {
			rv[idx] = n
			idx++
		}
	}

	for _, n := range A {
		if n%2 != 0 {
			rv[idx] = n
			idx++
		}
	}
	return rv
}

// 一次迭代，分别用两个 slice 保存奇数和偶数，最后合并
func divideAndMerge(A []int) []int {
	odd := make([]int, 0, len(A))
	even := make([]int, 0, len(A))
	for _, n := range A {
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}
	return append(even, odd...)
}
