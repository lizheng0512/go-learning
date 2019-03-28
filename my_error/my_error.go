package main

import (
	"fmt"
	"os"
)

type myError struct {
	msg  string
	code int
}

// 实现 error 接口
func (me *myError) Error() string {
	return me.msg
}

func main() {

	fmt.Println("args: ", os.Args)

	f := func(i int) (int, error) {
		if i <= 0 {
			return 0, &myError{msg: "i must > 0", code: 500}
		} else {
			return 10 / i, nil
		}
	}

	_, err := f(-2)
	if err != nil {
		if me, ok := err.(*myError); ok {
			fmt.Println(me.msg, me.code)
		}
	}

	// 使用 fmt.Errorf 来生成错误对象
	fmt.Println(fmt.Errorf("fmt.Errorf error"))

}
