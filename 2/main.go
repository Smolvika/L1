package main

import (
	"fmt"
	"sync"
)

func main() {
	var nums = []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, val := range nums {
		// Увеличиваем счетчик wg перед запуском горутины на 1
		wg.Add(1)
		val := val
		go func() {
			fmt.Println(val * val)
			// Уменьшаем счетчик wg по завершению работы горутины на 1
			wg.Done()
		}()
	}

	// Ждем пока счетчик wg не станет равным 0, т.е. завершения всех горутин
	wg.Wait()
}
