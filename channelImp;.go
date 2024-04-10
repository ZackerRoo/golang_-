package main

/* func main() {
// 其使用的场景是
// 停止信号监听
// 定时任务
// 生产者消费者模型解耦
// 控制并发数
/*
	主要四个主键比较重要

	1. gorouter之间传递数据的buf
	2. 当前正在接受和发送数据的下标
	3. 接受和发送数据被堵塞的队列
	4. 安全锁
*/
/*
	m := make(chan int) // 在无缓存的channel种发送和接受操作必须要不同的gorouter同时进行 无缓存意味着即发送即接受
	m <- 1              //这个操作会被堵塞直到有另外一个线程接受了这个数据
	<-m
	//因为 Go 语言中的通道默认是同步的，发送操作会一直等待，直到有另一个 goroutine 来接收这个值
}*/
/* func main() {

	var ch = make(chan int)
	var stopCh = make(chan struct{})
	// 多生产者
	for i := 1; i <= 100; i++ {
		go func(n int) {
			for {
				select {
				case ch <- n:
				case <-stopCh:
					return
				}
			}
		}(i)
	}

	// 单消费者
	go func() {
		for elem := range ch {
			fmt.Println(elem)
			if elem == 100 {

				close(stopCh)
				return
			}
		}
	}()

	// select {}
	for range ch {

	}
} */
