package main

import "fmt"

func main() {
	var a = [2]int{1, 1} //a为长度为2,int类型的数组
	var f *[2]int = &a   //*表示这个数组是个指针
	for i, v := range f {
		fmt.Println(i)
		fmt.Println(v)
		f[i] += 1
	}
	fmt.Printf("%v", a)
}
