package main

import (
	"fmt"
	"time"
)

func fibonacci(i int) int {
	if i <= 2 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	for i := 0; i < 100; i++ {
		go fmt.Println("goruntine: ", i)
	}
	fmt.Println("here")
	time.Sleep(10)
}
