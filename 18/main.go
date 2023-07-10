package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mx    sync.Mutex
}

func newCounter() *counter {
	return &counter{
		count: 0,
		mx:    sync.Mutex{},
	}
}

func (c *counter) add() {
	// Блокировка: доступ к данным может получить только одина горутина
	c.mx.Lock()
	// Разблокировка: предоставляет доступ к данным другим горутинам
	defer c.mx.Unlock()
	c.count++
}

func main() {
	wg := sync.WaitGroup{}
	c := newCounter()
	// Увеличиваем счетчик wg на 1000
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// Уменьшаем счетчик wg по завершению работы горутины на 1
			defer wg.Done()
			c.add()
		}()
	}
	// Ждем пока счетчик wg не станет равным 0, т.е. завершения всех горутин
	wg.Wait()
	fmt.Println(c.count)

}
