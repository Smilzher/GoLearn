package main

import (
	"GoLearn/imooc/crawler/engine"
	"GoLearn/imooc/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
