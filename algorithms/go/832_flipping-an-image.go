package main

import "fmt"

// 832. Flipping an Image
// https://leetcode.com/problems/flipping-an-image/
//
// Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.
//
// To flip an image horizontally means that each row of the image is reversed.
// For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].
//
// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
// For example, inverting [0, 1, 1] results in [1, 0, 0].

func main() {
	process([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	})

	process([][]int{
		{1, 1, 0, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 1, 0},
	})
}

func process(A [][]int) {
	showImage(A)
	rv := flipAndInvertImage(A)
	showImage(rv)
	fmt.Println("-----")
}

func showImage(A [][]int) {
	for _, row := range A {
		fmt.Println(row)
	}
	fmt.Println()
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		head, tail := 0, len(row)-1
		for head < tail {
			row[head], row[tail] = invert(row[tail]), invert(row[head])
			head++
			tail--
		}
		if head == tail {
			row[head] = invert(row[head])
		}
	}
	return A
}

func invert(n int) int {
	if n == 1 {
		return 0
	}
	return 1
}
