package main

import "fmt"

func main() {
	i := 2
	var arr = []int{2, 4, 6, 8, 10}
	delete3(arr, i)
	arr = arr[:len(arr)-1]
	fmt.Println(arr)

}

func delete1(nums []int, i int) {
	nums = append(nums[:i], nums[i+1:]...)
}

func delete2(arr []int, i int) {
	copy(arr[i:], arr[i+1:])
}

func delete3(arr []int, i int) {
	for ; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
}
