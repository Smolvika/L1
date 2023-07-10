package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{8, 7, 3, 9, 0, 1, 6, 5, 4, 2}
	sort.Ints(arr)

	fmt.Println(binarySearch(arr, 0))
}

func binarySearch(arr []int, num int) int {
	l, r := -1, len(arr)

	for r-l > 1 {
		// Ищем приблизительную середину
		m := (l + r) / 2

		if arr[m] < num {
			// Сдвигаем левую границу
			l = m
		} else if arr[m] > num {
			// Сдвигаем правую границу
			r = m
		} else {
			return m
		}
	}
	return -1
}
