package main

import "fmt"

// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/

func main() {
	for n := 0; n <= 30; n++ {
		fmt.Println(n, fib(n))
	}
}

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}

	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}
