package main

import "fmt"

// 804. Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	rv := uniqueMorseRepresentations(words)
	fmt.Println(rv)
}

var MorseTable = []string{
	// a,   b       c      d      e     f       g
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.",

	// h     i       j      k       l      m     n
	"....", "..", ".---", "-.-", ".-..", "--", "-.",

	// o     p       q       r      s     t
	"---", ".--.", "--.-", ".-.", "...", "-",

	// u     v       w       x       y      z
	"..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	resultSet := make(map[string]struct{}, 100)
	for _, w := range words {
		s := ""
		for _, c := range w {
			s += MorseTable[c-'a']
		}
		if _, ok := resultSet[s]; !ok {
			resultSet[s] = struct{}{}
		}
	}
	return len(resultSet)
}
