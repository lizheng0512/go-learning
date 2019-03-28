package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func fibnacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (ig intGen) Read(p []byte) (n int, err error) {
	next := ig()
	if next > 10000 {
		return 0, errors.New("done")
	}
	s := fmt.Sprintf("%d", next)
	fmt.Println("write")
	return strings.NewReader(s).Read(p)
}

func print(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Err(), scanner.Text())
	}
}

func main() {
	ig := fibnacci()
	print(ig)
}
