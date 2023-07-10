package main

import (
	"fmt"
	"sync"
)

func main() {
	var nums = []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	count := 0

	for _, val := range nums {
		// Увеличиваем счетчик wg перед запуском горутины на 1
		wg.Add(1)
		val := val
		go func() {
			// Блокировка: доступ к данным может получить только одина горутина
			mx.Lock()
			count += val * val
			// Разблокировка: предоставляет доступ к данным другим горутинам
			mx.Unlock()
			// Уменьшаем счетчик wg по завершению работы горутины на 1
			wg.Done()
		}()
	}
	// Ждем пока счетчик wg не станет равным 0, т.е. завершения всех горутин
	wg.Wait()
	fmt.Println(count)
}
