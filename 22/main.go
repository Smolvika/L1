package main

import (
	"fmt"
	"math/big"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	res := new(big.Int)
	a, ok1 := new(big.Int).SetString(str1, 10)
	b, ok2 := new(big.Int).SetString(str2, 10)
	if !ok1 || !ok2 {
		fmt.Println("Проверьте правильность введенных чисел")
		return
	}
	fmt.Printf("Cумма чисел: %d + %d = %d \n", a, b, res.Add(a, b))
	fmt.Printf("Разность чисел: %d - %d = %d \n", a, b, res.Sub(a, b))
	fmt.Printf("Произведение чисел: %d * %d = %d \n", a, b, res.Mul(a, b))
	fmt.Printf("Деление чисел: %d / %d = %d \n", a, b, res.Div(a, b))

}
