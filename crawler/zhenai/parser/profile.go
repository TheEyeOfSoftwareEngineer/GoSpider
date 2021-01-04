package parser

import (
	"GoSpider/crawler/engine"
	"GoSpider/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var locationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var heightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var hokouRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var styleRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-8b1eac0c>体型:([^<]+)</div>`)

//<div class="m-btn purple" data-v-8b1eac0c>36岁</div>
//<div class="m-btn purple" data-v-8b1eac0c>162cm</div>
//<div class="m-btn purple" data-v-8b1eac0c>54kg</div>
//<div class="m-btn pink" data-v-8b1eac0c>籍贯:陕西宝鸡</div>
//<div class="m-btn pink" data-v-8b1eac0c>体型:丰满</div>

func ParserProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	profile.Location = extractString(contents, locationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Style = extractString(contents, styleRe)

	result := engine.ParseResult{
		Items: []interface{} {profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}