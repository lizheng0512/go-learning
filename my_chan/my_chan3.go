package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
			time.Sleep(1e7)
		}
		close(ch)
	}()

	go func() {
		for i := range ch {
			fmt.Println("child: ", i)
		}
	}()

	for i := range ch {
		fmt.Println("main: ", i)
	}
}
