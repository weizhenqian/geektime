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
	}()
	//main执行到此，会等待c插入值
	<-c
}
