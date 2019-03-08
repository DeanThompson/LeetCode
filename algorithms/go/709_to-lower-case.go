package main

import "fmt"

// 709. To Lower Case
// https://leetcode.com/problems/to-lower-case/

func main() {
	for _, s := range []string{"Hello", "here", "LOVELY", "adf#2!@"} {
		fmt.Println(s, toLowerCase(s))
	}
}

func toLowerCase(str string) string {
	chars := []rune(str)
	delta := 'a' - 'A'
	for idx, c := range chars {
		if c >= 'A' && c <= 'Z' {
			chars[idx] = c + delta
		}
	}
	return string(chars)
}
