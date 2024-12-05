package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	fmt.Println("channels")

// 	ch := make(chan string)

// 	wg := sync.WaitGroup{}

// 	wg.Add(1)

// 	go func() {
// 		fmt.Println("Hello from pub function")
// 		ch <- "hello"
// 	}()
// 	go func() {
// 		fmt.Println("Hello from sub function")
// 		fmt.Println(<-ch)
// 		wg.Done()
// 	}()
// 	wg.Wait()

// }
