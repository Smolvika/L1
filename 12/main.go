package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})
	res := make([]string, 0)

	for _, val := range arr {
		// Заносим строку в map и в результирующий слайс, если ее еще нет в map
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			res = append(res, val)
		}
	}

	fmt.Println(res)
}
