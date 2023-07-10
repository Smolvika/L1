package main

import "fmt"

func main() {
	arr := []int{8, 7, 3, 9, 0, 1, 6, 5, 4, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, l, r int) {
	var i int
	if l < r {
		i = l

		for j := l; j < r; j++ {
			if arr[j] < arr[r] {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}

		arr[i], arr[r] = arr[r], arr[i]

		quickSort(arr, l, i-1)
		quickSort(arr, i+1, r)
	}
}
