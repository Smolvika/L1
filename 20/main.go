package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	res := reverse2(str)
	fmt.Println(res)
}

func reverse(s string) string {
	str := strings.Split(s, " ")

	var resStr string
	for _, val := range str {
		resStr = val + " " + resStr
	}

	return resStr
}

func reverse2(s string) string {
	str := strings.Split(s, " ")

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	return strings.Join(str, " ")
}
