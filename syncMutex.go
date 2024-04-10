package main

import "sync"

// 同步锁是用来保证多个共享数据的安全 在不同的线程之间 同步锁 有如下的特点阻塞 互斥 公平  需要防止死锁和饥饿出现
var lock sync.Mutex
var count int

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
}

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			increment()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	println(count)
// }
