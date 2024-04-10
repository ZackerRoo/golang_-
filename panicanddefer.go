package main

import "fmt"

/* func F() {
	defer func() {
		fmt.Println("b")
	}() // panic后面的defer 函数不会执行 由于这个在panic前面，所以会执行
	// 如果panic 有被recover捕获，那么defer函数全部会正常执行
	panic("a")
} */

// 子函数抛出的panic没有recover时，上层函数时，程序直接异常终止
/* func main() {
	defer func() {
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
} */

// b
// c
// panic: a

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}

/* func main() {
	defer func() {
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
} */

// 捕获异常: a
// b
// 继续执行
// c
