package main

// func main() {
// 	// s := "Hello, "

// 	// fmt.Println("Byte (uint8) view:")
// 	// for i := 0; i < len(s); i++ {
// 	// 	fmt.Printf("%x ", s[i])
// 	// }
// 	// fmt.Println("\nRune (int32) view:")
// 	// for i, r := range s {
// 	// 	fmt.Printf("%U ", r)
// 	// 	if i < len(s)-1 {
// 	// 		fmt.Print("| ")
// 	// 	}
// 	// }

// 	src := []int{1, 2, 3, 4, 5}

// 	var dst []*int
// 	for _, v := range src {
// 		dst = append(dst, &v) //为什么呢, 因为 for-range 中 循环变量的作用域的规则限制
// 		//假如取消append()后一行的注释，可以发现循环中v的变量内存地址是一样的，也可以解释为for range相当于
// 	}

// 	for _, v := range dst {
// 		println(*v)
// 	}

// 	fmt.Println(funcName("12"))
// }

func funcName(a interface{}) string { return a.(string) }
