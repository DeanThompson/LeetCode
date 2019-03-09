package main

import "fmt"

// 997. Find the Town Judge
// https://leetcode.com/problems/find-the-town-judge/
//
// In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.
//
// If the town judge exists, then:
//
// The town judge trusts nobody.
//   - Everybody (except for the town judge) trusts the town judge.
//   - There is exactly one person that satisfies properties 1 and 2.
//   - You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.
//
// If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.

func main() {
	process(1, [][]int{})
	process(2, [][]int{{1, 2}})
	process(3, [][]int{{1, 3}, {2, 3}})
	process(3, [][]int{{1, 3}, {2, 3}, {3, 1}})
	process(3, [][]int{{1, 2}, {2, 3}})
	process(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}})
}

func process(N int, trust [][]int) {
	judge := findJudge(N, trust)
	fmt.Println(N, trust, judge)
}

func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 {
		return 1
	}

	beTrustedTable := make(map[int]int, N)
	trustTable := make(map[int]struct{}, N)
	for _, pair := range trust {
		a, b := pair[0], pair[1]
		beTrustedTable[b] += 1
		trustTable[a] = struct{}{}
	}

	//fmt.Println(beTrustedTable, trustTable)
	judge := -1
	for n, persons := range beTrustedTable {
		//fmt.Println(n, persons, judge)
		if persons == N-1 {
			if judge != -1 {
				return -1
			}
			if _, ok := trustTable[n]; !ok {
				judge = n
				return judge
			}
		}
	}
	return judge
}
