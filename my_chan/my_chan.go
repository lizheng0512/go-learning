package main

import (
	"fmt"
	"time"
)

// 默认情况下, 通道是同步且无缓冲的, 无缓冲的通道没有空间来缓存数据, 所以在发送数据之前必须要接收者事先准备好
// 通道的发送和接收在对方准备好之前是阻塞的
func sendData(ch chan string) {
	ch <- "我"
	fmt.Println("sended 我")
	ch <- "爱"
	fmt.Println("sended 爱")
	ch <- "你"
	fmt.Println("sended 你")
	ch <- "中"
	fmt.Println("sended 中")
	ch <- "国"
	fmt.Println("sended 国")
}

func sendData2(ch chan string) {
	ch <- "哈"
	fmt.Println("sended 哈")
	ch <- "哈"
	fmt.Println("sended 哈")
}

func sendData3(ch chan string) {
	time.Sleep(5e9)
	ch <- "msg"
}

// 接收操作是阻塞的  如果通道中没有数据  接受者就阻塞了
func getData(ch chan string) {
	var data string
	for data = range ch {
		time.Sleep(1e9)
		fmt.Println(data)
	}
	fmt.Println("读取结束")
}
