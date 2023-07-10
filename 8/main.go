package main

import "fmt"

func main() {
	var num int64 = 12
	fmt.Printf("Number: %d\n in bytes: %08b\n", num, num)
	newNum := change(num, 4, 0)
	fmt.Printf("Number: %d\n in bytes: %08b\n", newNum, newNum)
}

func change(num int64, n int64, bit int64) int64 {
	if bit != 0 && bit != 1 {
		return num
	}

	if num&(1<<n) == bit<<n {
		return num
	} else if bit == 1 {
		return num | 1<<n
	} else {
		return num ^ (1 << n)
	}
}
