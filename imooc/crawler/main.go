package main

import (
	"github.com/lizheng0512/go-learning/imooc/crawler/engine"
	"github.com/lizheng0512/go-learning/imooc/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
