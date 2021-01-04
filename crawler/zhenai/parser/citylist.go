package parser

import (
	"GoSpider/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*data-v-1573aa7c>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	for _, subMatch := range m {
	//		fmt.Printf("%s ", subMatch)
	//	}
	//	fmt.Println()
	//}
	result := engine.ParseResult{}
	limit := 1

	for _, m := range matches {
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit --
		if limit == 0 {
			break
		}
	}

	return result
	//fmt.Printf("Matches found: %d\n", len(matches))

	//结果
	//City: 阿坝, URL: http://www.zhenai.com/zhenghun/aba
	//City: 阿克苏, URL: http://www.zhenai.com/zhenghun/akesu
	//City: 阿拉善盟, URL: http://www.zhenai.com/zhenghun/alashanmeng
	//City: 阿勒泰, URL: http://www.zhenai.com/zhenghun/aletai
	//City: 阿里, URL: http://www.zhenai.com/zhenghun/ali
	//City: 安徽, URL: http://www.zhenai.com/zhenghun/anhui
	// ...

}


