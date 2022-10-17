package parser

import (
	"GoLearn/imooc/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1) //有括号匹配出内容时，使用findAllSubmatch

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			}, //engine.NilParser,
		})
		//fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
	}
	return result
}
