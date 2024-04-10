package main

// var i int = 1
// var once sync.Once

/* func main() {
/*
	sync.Once 是一个机制能够确保在多个gorouter中某个操作只执行一次
*/
// once.Do()
// for j := 0; j < 10; j++ {
// 	go func() {
// 		once.Do(func() { //只会执行一次
// 			i++
// 			println(i)
// 		})
// 	}()
// }

/*
	atomic 原子操作
	指的是不会被中断的操作
	即要么全部成功要么全部失败
	这使得在多线程环境下不会出现数据竞争
*/

/* 	var a int32 = 1
for i := 0; i < 100; i++ {
	atomic.AddInt32(&a, 1)
}
fmt.Printf("a: %v\n", a) */

/*
	锁和原子操作的区别
	1. 原子操作是不会被中断的操作
	2. 原子操作只适用于简单的操作如修改数据 不适用与线程或是复杂数据结构的操作

*/
//}
