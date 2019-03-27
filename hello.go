package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v len()=%d cap()=%d", s, len(s), cap(s))
}

func main() {
	s1 := make([]int, 10)
	printSlice(s1)

}
