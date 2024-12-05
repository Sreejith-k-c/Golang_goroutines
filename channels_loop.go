package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels looping")
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("Data from channel > ", msg)
	}
}
