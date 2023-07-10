package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	stop1()
}

func stop1() {
	// Создаём дочерний контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			// Ждем закрытие канала
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("gorutine work")

			}

		}
	}()
	s := make([]byte, 1)
	os.Stdin.Read(s)
	// Вызываем функцию отмены
	cancel()
}

func stop2() {
	// Создаём дочерний контекст с тайм-аутом
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		for {
			select {
			// Ждем закрытие канала по истечению тайм-аута
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond)
				fmt.Println("gorutine work")

			}

		}
	}()

}

func stop3() {
	// Создаём дочерний контекст с дедлайном
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	go func() {
		for {
			select {
			// Ждем закрытие канала по наступлению дедлайна
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond)
				fmt.Println("gorutine work")

			}

		}
	}()
}

func stop4() {
	ch := make(chan struct{})
	go func() {
		for {
			select {
			// Ждем заптси в  канал
			case <-ch:
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("gorutine work")

			}

		}
	}()
	s := make([]byte, 1)
	os.Stdin.Read(s)
	// Записываем в  канал
	ch <- struct{}{}
}
