package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	arr := make([]int, 12)
	// Заполняем слайс
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	go func() {
		// Отправляем данные из слайса в канал ch1
		for _, val := range arr {
			ch1 <- val
		}
		close(ch1)
	}()

	go func() {
		// Читаем данные из канала ch1, возводим в квадрат и отпраялем в канал ch2
		for val := range ch1 {
			ch2 <- val * val
		}
		close(ch2)
	}()

	// Читаем данные из канала ch2
	for val := range ch2 {
		fmt.Println(val)
	}
}
