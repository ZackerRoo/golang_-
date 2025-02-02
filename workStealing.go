package main

//func main() {
/*
			workstealing 机制
		工作窃取是一种负载平衡策略。当一个P空闲时（即本地队列为空），
		它会尝试从其他P的本地队列中“窃取”一些Goroutine来执行。
		如果其他P的本地队列也都为空，它会尝试从全局队列中获取Goroutine。
	工作窃取有助于确保所有P都有工作可做，从而提高CPU利用率和程序性能。

	这样的好处是避免了线程饥饿的问题 提高gorouter的执行效率
*/

// runtime.GOMAXPROCS(2) // 启用两个 P

// var wg sync.WaitGroup
// wg.Add(10)

// for i := 0; i < 10; i++ {
// 	go func() {
// 		time.Sleep(time.Second) // 模拟 Goroutine 的耗时操作

// 		fmt.Println("Goroutine executed.")

// 		wg.Done()
// 	}()
// }

// wg.Wait()

/*
	handoff 和 workstealing 的区别
	1 主动性：Work Stealing 更多依赖于线程之间的主动性——空闲的线程主动去窃取其他线程的任务。
	而 Handoff 策略通常是一种被动的接受，即一个线程将任务显式地传递给另一个线程。
	2 目的和应用场景：Work Stealing 主要用于负载平衡，尽量让所有的线程都有工作可做
	，减少空闲。Handoff 则用于管理复杂的任务依赖和协作，特别是当任务之间有明确的前后依赖关系时。

*/

/*
	抢占式调度和协作式调度
	1. 抢占式调度：操作系统可以在任何时候中断一个线程，然后切换到另一个线程。
		抢占式采用了基于符号时间片的调度策略，即每个线程被分配一个时间片，当时间片用完时，操作系统会中断线程并切换到另一个线程。
	2. 协作式调度：线程只有在当前线程主动放弃CPU时，操作系统才会切换到另一个线程。
		比如说调用了time.Sleep()、调用了IO操作、调用了runtime.Gosched()等。让出时间片的方法实现了协作式调度。
*/

/*
	GMP调度过程中存在哪些阻塞
	1. 系统调用：比如文件读写、网络请求等，这些操作会导致线程阻塞，从而导致Goroutine被阻塞。
	2 通道操作：当通道为空时，从通道读取数据会阻塞；当通道满时，向通道写入数据会阻塞。
	3. 互斥锁也就是锁竞争：当一个Goroutine获取到互斥锁时，其他Goroutine就无法获取到锁，从而被阻塞。
	4. 垃圾回收：当垃圾回收器运行时，所有的Goroutine都会被暂停，从而导致阻塞。
*/

/*
	Sysmon的作用
	1. 处理垃圾回收：Sysmon会监控内存使用情况，当内存使用超过阈值时，会触发垃圾回收。
	2. 将长时间运行的Goroutine发出抢占式调度
	3. 将长时间没有处理的netpoll添加到全局队列
*/

//}
