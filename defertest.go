// package main

// import (
// 	"fmt"
// )

// func main() {

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println("fatal")
// 		}
// 	}()

// 	defer func() {
// 		panic("defer panic")
// 	}()

// 	panic("panic")
// 	// panic仅有最后一个可以被revover捕获。
// }

package main

import "fmt"

func function(index int, value int) int {

	fmt.Println(index)

	return index
}

// func main() {
// 	defer function(1, function(3, 0)) //那么这4个函数的先后执行顺序是什么呢？这里面有两个defer， 所以defer一共会压栈两次，先进栈1，后进栈2。 那么在压栈function1的时候，需要连同函数地址、函数形参一同进栈，那么为了得到function1的第二个参数的结果，所以就需要先执行function3将第二个参数算出，那么function3就被第一个执行。同理压栈function2，就需要执行function4算出function2第二个参数的值。然后函数结束，先出栈fuction2、再出栈function1.
// 	defer function(2, function(4, 0))
// }
