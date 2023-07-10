package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)

	for _, val := range arr {
		// Определяем ключ
		key := int(val) / 10 * 10
		// Записываем в map по ключу
		res[key] = append(res[key], val)
	}
	fmt.Println(res)
}
