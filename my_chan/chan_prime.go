package main

import (
	"fmt"
	"time"
)

// 生成从2开始的只读数字通道
func generate() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			time.Sleep(1e7)
			ch <- i
		}
	}()
	return ch
}

// 返回所有不能被 prime 整除的数
func filter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

// 生成质数序列通道
func sieve() <-chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
