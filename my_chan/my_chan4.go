package main

import (
	"fmt"
	"time"
)

func main() {
	// 实现一个线程池模式
	// 大小为10的线程池
	poolControl := make(chan uint8, 2)
	// 一共30个任务
	chResult := make(chan string, 30)
	for i := 0; i < 30; i++ {
		go func() {
			poolControl <- 0
			fmt.Println("mission start...")
			time.Sleep(3e9)
			<-poolControl
			chResult <- "mission done..."
		}()
	}

	for i := 0; i < 30; i++ {
		fmt.Println(<-chResult)
	}
}
