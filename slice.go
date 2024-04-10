package main

/*
	 func main() {
		_ = make([]string, 200)
		_ = make([]int, 200000)

		//  slice 截取x[]，x[low:high:max]，从下标low开始取元素，到high-1下标为止，容量为max-low 但是要保证在x的长度范围内 而不是cap(x)范围内

		list := new([]int) // new 返回的是指针所以返回了*[]int
		*list = append(*list, 1)
		*list = append(*list, 2)
		*list = append(*list, 3)
		*list = append(*list, 4)
		fmt.Println(*list)

		arr3 := [3]int{0: 3, 1: 4}
		fmt.Println(arr3)

		s := make([]int, 0)
		fmt.Println(BeforeAppend(s))
		fmt.Println(AfterAppend(s))

}
*/
func BeforeAppend(s []int) []int {
	s = append(s, 1)
	s = []int{1, 2, 3}
	return s
}

func AfterAppend(s []int) []int {
	s = []int{1, 2, 3}
	s = append(s, 1)
	return s
}
