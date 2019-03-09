package main

import "fmt"

// 933. Number of Recent Calls
// https://leetcode.com/problems/number-of-recent-calls/
//
// 思路：最朴素的想法就是有个数组保存所有 t，传入新的 t 时，
// 从数组里找出 [t-3000, t] 区间里的元素数量。
// 一种思路就是从头开始迭代，直到遇到第一个 >= t-3000，取下标为 idx, 返回 length - idx
// 但实际上保留所有 t 是没有必要的，考虑到 t 是递增的，其实只要保留 [t-3000, t] 就可以了
// 下次输入新的 t 就可以减少迭代次数。 这个其实就是一个双向列表。

func main() {
	rc := Constructor()
	for _, t := range []int{1, 100, 3001, 3002} {
		c := rc.Ping(t)
		fmt.Println(t, c)
	}
}

type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{times: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)
	idx := 0
	for {
		if this.times[idx] >= t-3000 {
			break
		}
		idx++
	}
	this.times = this.times[idx:]
	return len(this.times)
}
