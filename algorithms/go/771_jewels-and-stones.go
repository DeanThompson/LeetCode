package main

import "fmt"

// 771. Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/

func main() {
	j := "aA"
	s := "aAAbbbb"
	rv := numJewelsInStones(j, s)
	fmt.Println(rv)
}

func numJewelsInStones(J string, S string) int {
	return solutionWithMap(J, S)
}

func solutionWithDoubleLoop(J string, S string) int {
	rv := 0
	for _, c := range J {
		for _, s := range S {
			if c == s {
				rv += 1
			}
		}
	}
	return rv
}

func solutionWithMap(J string, S string) int {
	counters := make(map[rune]int, len(S))
	for _, s := range S {
		counters[s] += 1
	}

	rv := 0
	for _, c := range J {
		rv += counters[c]
	}
	return rv
}
