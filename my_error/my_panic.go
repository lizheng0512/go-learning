package main

import (
	"fmt"
	"github.com/gogap/logrus"
	"reflect"
	"strconv"
)

func main() {

	// 在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给
	// 接收到 panic 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用
	// 传给 panic 的值报告错误情况：这个终止过程就是 panicking。

	fmt.Println("start...")

	ch := make(chan map[string]interface{}, 1)

	occurPanic(ch)

	fmt.Println(<-ch)
	fmt.Println("end")

}

func occurPanic(ch chan map[string]interface{}) {
	// panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止
	arr := make([]int, 2)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(reflect.TypeOf(err))
			logrus.Errorf("recover: arr: %v, msg: %s", arr, err)
			ch <- make(map[string]interface{}, 0)
		}
	}()

	panic(&customError{msg: "123", code: 500})
	fmt.Println(arr[2])

	ch <- map[string]interface{}{"1": 1}
}

type customError struct {
	msg  string
	code int
}

// 实现 error 接口
func (me *customError) Error() string {
	return me.msg + strconv.Itoa(me.code)
}
