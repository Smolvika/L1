package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
}

func createHugeString(size int) string {
	rand.NewSource(time.Now().UnixNano())

	builder := strings.Builder{}

	for i := 0; i < size; i++ {
		builder.WriteRune('a' + rune(rand.Intn('z'-'a'+1)))
	}
	return builder.String()
}

func main() {
	someFunc()

}
