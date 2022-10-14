package parser

import (
	"GoLearn/imooc/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1) //有括号匹配出内容时，使用findAllSubmatch

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
	}
	return result
}
