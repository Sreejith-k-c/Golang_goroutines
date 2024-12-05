package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func taskA(wg *sync.WaitGroup) {
// 	fmt.Println("Task A")
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("Task A Done")
// 	wg.Done()
// }

// func main() {
// 	fmt.Println("concunnrency")

// 	var wg sync.WaitGroup

// 	wg.Add(1)

// 	go taskA(&wg)

// 	wg.Wait()

// 	// time.Sleep(time.Second * 2)

// 	fmt.Println("main exit")
// }
