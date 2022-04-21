/*
基于 Channel 编写一个简单的单线程生产者消费者模型：

队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
package main

import (
	"fmt"
	"time"
)

//定义两个定时器，一个快一个慢，用于测试阻塞功能
var fastTicker = time.NewTicker(1 * time.Second)
var slowTicker = time.NewTicker(10 * time.Second)

//定义chan
var queue = make(chan int, 10)

//消费者功能
func consumer() {
	for {
		for _ = range slowTicker.C {
			i, ok := <-queue
			if !ok {
				fmt.Println("producer is block")
			}
			fmt.Println(i)
		}
	}
}

//生产者功能
func producer() {
	for _ = range fastTicker.C {
		fmt.Println("input chan int")
		queue <- 1
	}
}

//函数主程序
func main() {
	go producer()
	consumer()
}
