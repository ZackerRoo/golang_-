package main

// func main() {
/*
	go线程共享相同的地址空间 内存和栈堆
	这样的化变量也是共享的
*/

/*

	GMP 模型
	G: Gorouter
	M: Machine
	P: Processor
	每个Gorouter都会被分配到一个P上
	每个P上都会有一个M
	每个M都会有一个线程
	这样的话就可以实现多线程并发

	M 是操作系统线程 每次一个M只会执行一个Gorouter 当有Gorouter被阻塞
	会将M交给其他Gorouter执行

	GMP的执行步骤如下
	1. 创建一个Gorouter
	2. 将Gorouter放到P的队列中
	3. P从队列中取出一个Gorouter执行
	4. 当Gorouter阻塞时 将M交给其他Gorouter执行
	5. 当Gorouter执行完毕时 会被销毁


*/
// }
//
