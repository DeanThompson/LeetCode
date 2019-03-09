package main

import "fmt"

// 657. Robot Return to Origin
// https://leetcode.com/problems/robot-return-to-origin/

func main() {
	for _, moves := range []string{"UD", "LL"} {
		fmt.Println(moves, judgeCircle(moves))
	}
}

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		switch c {
		case 'U':
			y += 1
		case 'D':
			y -= 1
		case 'L':
			x -= 1
		case 'R':
			x += 1
		}
	}
	return x == 0 && y == 0
}
