package main

/* func main() {
// 一般而言要实现线程安全的方法有如下机制
/*
	Mutex锁
	读写锁
	原子操作
	sync.once
	sync.atomic
	channel
*/

/*
	而slice 不是线程安全的主要原因是
	其底层数组并没有上面的加锁机制，所以在并发的情况下，可能会出现多个goroutine同时操作一个slice的情况
*/
//}
