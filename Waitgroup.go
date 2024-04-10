package main

/*
func main() {
	/*
		waitgropu 用于等待一组gorouter的结束
		其实Waitgroup的底层还是cond来实现的
		Waitgroup的Add()方法用于添加等待的gorouter数量
		Waitgroup的Done()方法用于减少等待的gorouter数量
		Waitgroup的Wait()方法用于等待所有的gorouter结束
		下面是一个示例代码

*/

// var wg sync.WaitGroup
// for i := 0; i < 5; i++ {
// 	wg.Add(1)
// 	go func(i int) {
// 		fmt.Printf("Gorouter %d is now running \n", i)
// 		wg.Done()
// 	}(i)
// }
// wg.Wait()
//}
