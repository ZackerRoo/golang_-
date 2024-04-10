package main

/* func main() {
/*
	Mutex 的几种状态MutexZeroValue, MutexLocked, MutexWoken, MutexStarving
	1. MutexZeroValue: 未初始化的Mutex
	2. MutexLocked: 互斥锁被锁定
	3. MutexWoken: 互斥锁被唤醒
	4. MutexStarving: 互斥锁处于饥饿状态
	还有一个waiter字段，用于记录等待锁的goroutine数量

	Mutex的两种模式一个是正常模式一个是饥饿状态
	正常模式是公平锁 所有等待的gorouter按照FIFO的顺序获取锁
	唤醒的gorouter不会直接获得锁而是会和新来的gorouter一起竞争 新来的gorouter有可能会获得锁

	饥饿模式 是非公平锁为了解决gorouter的长尾问题
	会直接把锁交给队列中排在最前面的gorouter 同时饥饿状态下新进来的gorouter直接进入队列 不会抢锁和自旋

	Mutex今天自旋的条件
	1 正常模式下的Mutex 会自旋 锁被占用并且Mutex不处于饥饿状态
	2 自旋次数小于4次
	3 cpu 个数大于1
	4 有空闲的p
	5 当前gorouter所挂载的p上没有其他的gorouter

	RWMutex 读写锁
	有两个重要细节
	1 优先唤醒写gorouter
	2 读锁等待队列和写锁等待队列是分开的

	读锁被解锁后 会唤醒写锁等待队列中的gorouter
	写锁被解锁后 会唤醒读锁等待队列中的gorouter

*/

/* var mutex sync.Mutex
cond := sync.NewCond(&mutex)
ready := false
for i := 0; i < 5; i++ {
	go func(i int) {
		mutex.Lock()
		for !ready {
			cond.Wait()
		}
		println(i)
		fmt.Printf("Gorouter %d is now running \n", i)
		mutex.Unlock()
	}(i)
}
time.Sleep(time.Second)
mutex.Lock()
ready = true
cond.Signal() //Signal()方法只会唤醒一个goroutine Broadcast()会唤醒所有的goroutine
mutex.Unlock()
time.Sleep(time.Second)

/*
	cond.Wait()
	这个方法是释放锁并且挂起当前gorouter
	直到有线程调用Signal()或者Broadcast()方法唤醒当前gorouter
	wait需要和lock配合使用否则会panic
*/
//}*/
