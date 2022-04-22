package main

import (
	"fmt"
)

func main() {
	//创建channel,值是布尔类型
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!!")
		//匿名函数完成时，插入true到channel
		c <- true
		//关闭c
		close(c)
	}()
	//此处循环c进行迭代，close(c)是关闭操作
	for v := range c {
		fmt.Println(v)
	}
}
