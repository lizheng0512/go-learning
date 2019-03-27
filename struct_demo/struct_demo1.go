package main

import "fmt"

type myStruct struct {
	value string
}

func (ms myStruct) print() {
	fmt.Println(ms.value)
}

func (ms *myStruct) printByPtr() {
	fmt.Println(ms.value)
}

func (ms myStruct) String() string {
	return ms.value
}

type myInterface interface {
	Get() string
}

func (ms *myStruct) Get() string {
	return ms.value
}

func main() {
	var ms1 fmt.Stringer = &myStruct{value: "123"}
	//ms1.print()
	//ms1.printByPtr()
	fmt.Println(ms1.String())
	var ms2 myInterface = &myStruct{value: "456"}
	fmt.Println(ms2.Get())
}
