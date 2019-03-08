package main

import (
	"fmt"
	"strings"
)

// 929. Unique Email Addresses
// https://leetcode.com/problems/unique-email-addresses/

func main() {
	emails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}

	rv := numUniqueEmails(emails)
	fmt.Println(rv)
}

func numUniqueEmails(emails []string) int {
	uniqEmails := make(map[string]struct{}, len(emails))
	for _, s := range emails {
		e := cleanEmailUsingStd(s)
		uniqEmails[e] = struct{}{}
	}

	return len(uniqEmails)
}

// 遍历字符串，保留 + 之前的非 . 字符
func cleanEmail(s string) string {
	chars := make([]rune, 0, 100)
	cursor := 0
	for idx, c := range s {
		cursor = idx
		if c == '+' || c == '@' {
			break
		}
		if c != '.' {
			chars = append(chars, c)
		}
	}

	for {
		if s[cursor] == '@' {
			break
		}
		cursor++
	}
	chars = append(chars, []rune(s[cursor:])...)
	return string(chars)
}

// 使用标准库的 Split 和 Replace 来处理
func cleanEmailUsingStd(s string) string {
	parts := strings.Split(s, "@")
	name := strings.Replace(parts[0], ".", "", -1)
	name = strings.Split(name, "+")[0]
	return fmt.Sprintf("%s@%s", name, parts[1])
}
