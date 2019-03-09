package main

import "fmt"

// 1002. Find Common Characters
// https://leetcode.com/problems/find-common-characters/
//
// 思路：构建哈希表，统计每个字符在所有字符串中出现的次数（大于 0），取最小的次数。
// 如何统计次数：为了减少循环次数，用数组保存每次出现的下标，最后统计数组长度
// 优化：因为要在每个字符串都出现，可以第一个字符串作为基准，只处理在该字符串中出现的字符

func main() {
	for _, A := range [][]string{
		{"bella", "label", "roller"},
		{"cool", "lock", "cook"},
		{"dadaabaa", "bdaaabcc", "abccddbb", "bbaacdba", "ababbbab", "ccddbbba", "bbdabbda", "bdabaacb"},
	} {

		fmt.Println(A, commonChars(A))
	}
}

func commonChars(A []string) []string {
	total := len(A)

	// 用一个哈希表保存每个字符在所有字符串中出现的次数
	table := make(map[rune][][]int, len(A[0]))
	for i, c := range A[0] {
		indexTable, ok := table[c]
		if !ok {
			indexTable = make([][]int, total)
			indexTable[0] = make([]int, 0, len(A[0]))
			table[c] = indexTable
		}
		indexTable[0] = append(indexTable[0], i)
	}

	for idx := 1; idx < total; idx++ {
		s := A[idx]
		for i, c := range s {
			if indexTable, ok := table[c]; ok {
				indexTable[idx] = append(indexTable[idx], i)
			}
		}
	}

	counter := make(map[rune]int, total)
	for c, indexTable := range table {
		if len(indexTable) == total {
			counter[c] = minSize(indexTable)
		}
	}
	result := make([]string, 0)
	for c, count := range counter {
		for i := 0; i < count; i++ {
			result = append(result, string(c))
		}
	}
	return result
}

func minSize(s [][]int) int {
	value := len(s[0])
	for _, x := range s {
		if len(x) < value {
			value = len(x)
		}
	}
	return value
}
