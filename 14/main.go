package main

import "fmt"

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case bool:
		fmt.Println("type is bool")
	case chan string:
		fmt.Println("type is chan int")
	default:
		fmt.Println("type is unknown")
	}
}

func main() {
	ch1 := make(chan string)

	checkType(0)
	checkType("abc")
	checkType(false)
	checkType(ch1)
	checkType(12.3)
}
