package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	numChan := make(chan int, 10)
	//done := make(chan bool)
	//go numGen(0, 10, numChan)
	//go numEchoRange(numChan, done)
	//
	//<-done

	go write(numChan)
	//go read(numChan, done)

	for i := 0; i < 10; i++ {
		num, ok := <-numChan
		fmt.Println(num, ok)
	}
	close(numChan)

	//<-done
}

// chan<- 表示只能往里插的通道
// <-chan 表示只能往外读的通道
// 两者都可以接收双向通道, 单在方法块里只能执行相应操作
func numGen(start, count int, out chan<- int) {
	defer close(out)
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
}

func numEchoRange(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println(num)
	}
	done <- true
}

func write(ch chan<- int) {
	// 写完了之后, 就要调用 close, 否则会造成一致循环读的接受者发生死锁
	//defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func read(ch <-chan int, resCh chan<- bool) {
	// 一致循环读, 如果通道已经关闭, 则会读出零值
	//for {
	//	i, ok := <-ch
	//	if ok {
	//		fmt.Println(i)
	//	} else {
	//		fmt.Println("读完了")
	//		resCh <- true
	//		return
	//	}
	//}
	// 这种循环方式, 在读取完通道中的数据后, 通道关闭之后, 会自动退出循环
	for i := range ch {
		fmt.Println(i)
	}
	resCh <- true
}
