package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {
	t1 := 1
	mySleep1(t1)
	//t2 := 1 * time.Second
	//mySleep2(t2)
	fmt.Printf("время истекло")
}

func mySleep1(t int) {
	<-time.After(time.Duration(t) * time.Second)
}

func mySleep2(t time.Duration) {
	<-time.After(t)
}
