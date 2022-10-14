package parser

import (
	"GoLearn/imooc/crawler/engine"
	"GoLearn/imooc/crawler/model"
	"regexp"
	"strings"
)

var ageRe = regexp.MustCompile(`<div.*class="des f-cl">[^>]*|(.*)|.*|.*|.*|[^<]*<a.*>.*</a></div>`)
var educationRe = regexp.MustCompile(`<div.*class="des f-cl">[^>]*|.*|(.*)|.*|.*|[^<]*<a.*>.*</a></div>`)
var marriageRe = regexp.MustCompile(`<div.*class="des f-cl">[^>]*|.*|.*|(.*)|.*|[^<]*<a.*>.*</a></div>`)
var heightRe = regexp.MustCompile(`<div.*class="des f-cl">[^>]*|.*|.*|(.*)|.*|[^<]*<a.*>.*</a></div>`)
var incomeRe = regexp.MustCompile(`<div.*class="des f-cl">[^>]*|.*|.*|.*|(.*)|[^<]*<a.*>.*</a></div>`)

//const profileRe = `<div.*class="des f-cl">([^>]*)|(.*)|(.*)|(.*)|(.*)|([^<]*)<a.*>.*</a></div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Age = extractString(contents, ageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Height = extractString(contents, heightRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents) //查找值之中的一个
	if len(match) >= 2 {
		return strings.Trim(string(match[1]), " ")
	} else {
		return ""
	}
}
