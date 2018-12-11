package main

import (
	"fmt"
)

func main() {
	natural := make(chan int)
	square := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			natural <- i
			// time.Sleep(1)
		}

	}()

	go func() {
		for {
			i := <-natural
			square <- i * i
		}
	}()

	for {
		y := <-square
		fmt.Println(y)

	}
}
