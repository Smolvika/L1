package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	n := 10
	// Создаём дочерний контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for {
			select {
			// Ждем закрытие канала
			case <-ctx.Done():
				return
			case ch <- rand.Int():
			}
		}
		close(ch)
	}()

	for i := 0; i < n; i++ {
		i := i
		go func() {
			// Увеличиваем счетчик wg на 1
			wg.Add(1)
			for v := range ch {
				fmt.Printf("воркер: %d  значение:%d\n", i, v)
			}
			// Уменьшаем счетчик wg  на 1
			wg.Done()
		}()
	}
	// Ждем пока счетчик wg не станет равным 0
	wg.Wait()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// Вызываем функцию отмены
	cancel()
}
