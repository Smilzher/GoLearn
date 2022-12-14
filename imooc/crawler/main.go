package main

import (
	"GoLearn/imooc/crawler/engine"
	"GoLearn/imooc/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
