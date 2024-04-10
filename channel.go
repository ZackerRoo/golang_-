package main

import (
	"fmt"
	"time"
)

/*func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// ch := make(chan int, 5)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}

	// 	wg.Done()
	// }()

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-ch)
	// 	}
	// 	wg.Done()
	// }()
	// wg.Wait()
	// close(ch)

	ch := make(chan int, 5)
	ch <- 6
	close(ch)
	//ch <- 2 // 一个channel 被关闭之后 不能再向里面写数据 但是可以读数据
	fmt.Println(<-ch)
	fmt.Println(<-ch) //而当channel空了之后 仍然可以读取但是读取的数据是0值

	ch1 := make(chan int, 5)
	go producer(ch1)
	consumer(ch1)

	可以使用channel 来控制gorouter的执行顺序

	channel 在如下的情况下会发生死锁
	1. 发送者和接收者不一致
	2. 单向channel使用不恰当
	3. 循环等待
	4. channel为空或者满了
	5. 未关闭channel
	6. 未接收channel数据


}*/

func producer(ch chan<- int) { //参数声明了一个单向的channel 只能写数据
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Consumer: %d\n", v)
		time.Sleep(time.Second)
	}
}
