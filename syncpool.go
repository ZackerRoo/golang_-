package main

/* func main() {
/*
	sync.Pool
	主要用于存储临时对象 以减少内存分配的次数和垃圾回收的压力
	内存重用的机制 也就是对象复用
*/

/* 	var pool = sync.Pool{
	New: func() any {
		return new(Object)
	},
} */

/* 	Object1 := pool.Get().(*Object)
	fmt.Printf("Object1: %v\n", &Object1)

	pool.Put(Object1)

	Object2 := pool.Get().(*Object)
	fmt.Printf("Object2: %v\n", &Object2)

}

type Object struct {
}
*/
