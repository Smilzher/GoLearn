package engine

import (
	"GoLearn/imooc/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request //相当于请求队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0] //取出第一个请求
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:error "+"fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body) //解析内容
		//parseResult.Requests...表示slice铺展开
		//将解析后的请求加入slice队列中
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
