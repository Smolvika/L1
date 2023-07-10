package main

import (
	"fmt"
	"sync"
)

type cache struct {
	m  map[int]int
	mx sync.Mutex
	wg sync.WaitGroup
}

func newCache() *cache {
	return &cache{
		m:  make(map[int]int),
		mx: sync.Mutex{},
		wg: sync.WaitGroup{},
	}
}

func (c *cache) add() {
	for i := 0; i < 20; i++ {
		// Увеличиваем счетчик wg перед запуском горутины на 1
		c.wg.Add(1)
		i := i
		go func() {
			// Уменьшаем счетчик wg по завершению работы горутины на 1
			defer c.wg.Done()
			// Блокировка: доступ к данным может получить только одина горутина
			c.mx.Lock()
			c.m[i] = i * i
			// Разблокировка: предоставляет доступ к данным другим горутинам
			c.mx.Unlock()
		}()
	}
	// Ждем пока счетчик wg не станет равным 0, т.е. завершения всех горутин
	c.wg.Wait()
}
func main() {
	cache := newCache()
	cache.add()
	fmt.Println(cache.m)

}
