package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		// go
		for {
			select {
			case <-stop:
				fmt.Println("stop channel")
				return //不写return，routing就泄露
			default:
				fmt.Println("Printing")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("stopping")
	stop <- true
	time.Sleep(5 * time.Second)
}
