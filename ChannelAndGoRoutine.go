package main

/* func main() {
	// limit := 26

	// numChan := make(chan int, 1)
	// charChan := make(chan int, 1)
	// mainChan := make(chan int, 1)
	// charChan <- 1

	// go func() {
	// 	for i := 0; i < limit; i++ {
	// 		<-charChan
	// 		fmt.Printf("%c\n", 'a'+i)
	// 		numChan <- 1

	// 	}
	// }()
	// go func() {
	// 	for i := 0; i < limit; i++ {
	// 		<-numChan
	// 		fmt.Println(i)
	// 		charChan <- 1

	// 	}
	// 	mainChan <- 1
	// }()
	// <-mainChan //利用mainChan来阻塞主线程，等待两个协程执行完毕
	// close(charChan)
	// close(numChan)
	// close(mainChan)

	select {} // 没有可以运行的协程，主线程阻塞
}
*/
