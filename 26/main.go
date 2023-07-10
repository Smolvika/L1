package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abCdefAaf"
	fmt.Println(unique(strings.ToLower(str)))
}

func unique(str string) bool {
	m := make(map[rune]struct{})
	for _, val := range str {
		if _, ok := m[val]; ok {
			return false
		}
		m[val] = struct{}{}
	}
	return true
}
