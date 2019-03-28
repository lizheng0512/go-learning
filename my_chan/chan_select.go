package main

import (
	"fmt"
	"time"
)

func main() {
	//
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go f1(ch1)
	//go f2(ch2)
	//
	//for {
	//	time.Sleep(1e7)
	//	// 当 case 字句的通道阻塞时, 会一直执行 default 子句
	//	select {
	//	case i1 := <-ch1:
	//		fmt.Println(i1)
	//	case i2 := <-ch2:
	//		fmt.Println(i2)
	//	default:
	//		fmt.Println("nothing")
	//	}
	//}

	// 简单超时实现
	// 缓冲大小设置为 1 是必要的，可以避免协程死锁以及确保超时的通道可以被垃圾回收
	//ch3 := make(chan int, 1)
	//go func() {
	//	time.Sleep(3e9)
	//	ch3 <- 1
	//}()
	//select {
	//case <-ch3:
	//	fmt.Println("success")
	//case <-time.After(5e9):
	//	fmt.Println("timeout")
	//default:
	//	fmt.Println("default")
	//}

	ch4 := make(chan string, 1)
	for i := 0; i < 10; i++ {
		//go func() {
		select {
		case ch4 <- f3():
		default:
			fmt.Println("default")
		}
		//}()
	}
	fmt.Println(<-ch4)
	time.Sleep(2e9)
}

func f1(ch chan<- int) {
	time.Sleep(4e9)
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func f2(ch chan<- int) {
	time.Sleep(5e9)
	for i := 0; i < 10; i++ {
		ch <- i + 10
	}
}

func f3() string {
	fmt.Println("f3")
	time.Sleep(1e9)
	return "hello"
}
