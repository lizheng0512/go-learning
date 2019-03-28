package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is liz0581@163.com
email2 is abc@163.com
email3 is 123@qq.com
email4 is ddd@123.com.cn
`

func main() {
	re := regexp.MustCompile(
		"([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\\.([a-zA-Z]+)")
	match := re.FindAllString(text, -1)
	// 同时找到所有的字串
	sub := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	fmt.Println(sub)
}
