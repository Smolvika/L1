package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10 * time.Second
	// Создаём дочерний контекст с тайм-аутом
	ctx, _ := context.WithTimeout(context.Background(), n)
	ch := make(chan int)

	go func() {
		for {
			select {
			// Ждем закрытие канала по истечению тайм-аута
			case <-ctx.Done():
				return
			case ch <- rand.Int():
			}
		}
		close(ch)

	}()
	// Читаем из канала
	for v := range ch {
		fmt.Printf("значение:%d\n", v)
	}

}
