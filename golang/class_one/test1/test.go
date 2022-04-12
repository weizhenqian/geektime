package main

import "fmt"

type Test struct {
	name string
}

func main() {
	test := Test{"test"}
	fmt.Println(test)
	//结果{test}

	testa := &Test{"test"}
	fmt.Println(testa)
	//fmt.Println(testa)
	//结果 &{test}

	testc := &Test{"test"}
	fmt.Println(*testc)
	//结果 {test}

	testd := &Test{"test"}
	fmt.Println(&testd)
	//结果 0xc000006030
}
