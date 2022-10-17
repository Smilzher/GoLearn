package parser

import (
	"GoLearn/imooc/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1) //有括号匹配出内容时，使用findAllSubmatch

	result := engine.ParseResult{}
	limit := 2
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
