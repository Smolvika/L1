package main

import "fmt"

func main() {
	arr1 := []int{10, -12, 3, 0, 9, -11, 7}
	arr2 := []int{11, -23, 10, -4, 9, 6, 7, -8, 0}
	// Объединяем слайсы
	arr3 := append(arr1, arr2...)

	m := make(map[int]int)
	// Проходим по объединенному слайсу и записываем в map
	for _, val := range arr3 {
		m[val]++
	}

	res := make([]int, 0)

	for key, val := range m {
		// Если число встречалось больше двух раз, то заносим в результирующий слайс
		if val > 1 {
			res = append(res, key)
		}
	}

	fmt.Println(res)
}
