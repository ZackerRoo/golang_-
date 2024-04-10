package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Work Done")
			return
		default:
			fmt.Println("Working")
			time.Sleep(1 * time.Second)

		}
	}
}

// func main() {
// 	ch := make(chan bool)
// 	go worker(ch)
// 	time.Sleep(5 * time.Second)
// 	ch <- true
// 	time.Sleep(1 * time.Second) // 防止主线程先结束了
// 	fmt.Println("Main Done")

// }
