package parser

import (
	"GoLearn/imooc/crawler/engine"
	"GoLearn/imooc/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)

//const profileRe = `<div.*class="des f-cl">([^>]*)|(.*)|(.*)|(.*)|(.*)|([^<]*)<a.*>.*</a></div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
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
