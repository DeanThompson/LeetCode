package main

import "fmt"

// https://leetcode.com/problems/sum-of-even-numbers-after-queries/
// 985. Sum of Even Numbers After Queries
//
// We have an array A of integers, and an array queries of queries.
//
// For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].
// Then, the answer to the i-th query is the sum of the even values of A.
//
// (Here, the given index = queries[i][1] is a 0-based index, and each query permanently
// modifies the array A.)
//
// Return the answer to all queries.  Your answer array should have answer[i] as the answer
// to the i-th query.

func main() {
	A := []int{1, 2, 3, 4}
	queries := [][]int{
		{1, 0},
		{-3, 1},
		{-4, 0},
		{2, 3},
	}
	result := sumEvenAfterQueries(A, queries)
	fmt.Println(A, queries, result)
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}

	result := make([]int, 0, len(queries))
	for _, q := range queries {
		val, idx := q[0], q[1]
		oldVal := A[idx]
		wasEven := A[idx]%2 == 0
		A[idx] += val
		isEven := A[idx]%2 == 0

		// 1. 之前是奇数，现在是偶数：加上结果
		// 2. 之前是奇数，现在还是奇数：不变
		// 3. 之前是偶数，现在是奇数：减去原来的值
		// 4. 之前是偶数，现在还是偶数：加上查询的值
		if wasEven && isEven {
			sum += val
		} else if wasEven && !isEven {
			sum -= oldVal
		} else if !wasEven && isEven {
			sum += A[idx]
		}
		result = append(result, sum)
	}
	return result
}
