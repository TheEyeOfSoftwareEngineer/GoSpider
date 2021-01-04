package main

import (
	"GoSpider/crawler/engine"
	"GoSpider/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}


// IDEA函数返回值自动补全 option+command+v

// 元素过滤
	// CSS选择器
	// xpath 能力没有CSS选择器强
	// 正则表达式

	