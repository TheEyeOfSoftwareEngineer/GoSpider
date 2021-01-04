package parser

import (
	"GoSpider/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`
//<a href="http://album.zhenai.com/u/1385132990" target="_blank">飞花落砚</a>
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	//fmt.Println("Matches: ", len(matches))

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, "User " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, string(m[2]))
			},
		})

	}

	return result
}