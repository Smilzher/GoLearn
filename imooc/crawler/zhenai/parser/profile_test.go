package parser

import (
	"GoLearn/imooc/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "金芝妤")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile) //类型断言？

	expected := model.Profile{
		Age:       "31岁",
		Education: "中专",
		Marriage:  "离异",
		Height:    "168cm",
		Income:    "12001-20000元",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
