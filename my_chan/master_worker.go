package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 使用通道进行同步: 使用一个通道接收要处理的任务, 一个通道接收处理完成的任务及其结果, worker 在协程中启动
func main() {
	pending, done := make(chan string), make(chan map[string]interface{})
	go sendWork(pending)
	for i := 0; i < 10; i++ {
		go worker(pending, done)
	}
	consumeWork(done)
}

func sendWork(ch chan string) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- strconv.Itoa(i)
	}
}

func worker(in chan string, out chan map[string]interface{}) {
	//defer close(out)
	for i := range in {
		randomInt := rand.Int63n(10)
		time.Sleep(time.Duration(randomInt) * 1e9)
		out <- map[string]interface{}{i: randomInt}
	}
}

func consumeWork(in chan map[string]interface{}) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-in)
	}
}
