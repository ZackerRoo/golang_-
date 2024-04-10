package main

// func main() {
// 	// 反射机制是一种在项目运行时候检查变量的机制，可以获取变量的类型信息，值信息，方法信息等等
// 	// 尽量 避免在项目中使用反射，因为反射会降低代码的可读性，可维护性，性能 其次反射是一个比较危险的操作
// 	// 其次尽量使用静态的检测工具，比如 go vet, go lint, go fmt 等等

// 	var x interface{}
// 	x = 100
// 	t := reflect.TypeOf(x)
// 	v := reflect.ValueOf(x)
// 	fmt.Printf("t: %v\n", t)
// 	fmt.Printf("v: %v\n", v)
// }
