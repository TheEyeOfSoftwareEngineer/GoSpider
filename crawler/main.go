package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	//http://www.zhenai.com/zhenghun
	response, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", response.StatusCode)
		return
	}

	e := determineEncoding(response.Body)
	//encoding转换
	newReader := transform.NewReader(response.Body,
						e.NewDecoder())
	all, err := ioutil.ReadAll(newReader)

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	printCityList(all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	//猜测获取HTML的encoding
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*data-v-1573aa7c>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	for _, subMatch := range m {
	//		fmt.Printf("%s ", subMatch)
	//	}
	//	fmt.Println()
	//}
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}

	fmt.Printf("Matches found: %d\n", len(matches))

	//结果
	//City: 阿坝, URL: http://www.zhenai.com/zhenghun/aba
	//City: 阿克苏, URL: http://www.zhenai.com/zhenghun/akesu
	//City: 阿拉善盟, URL: http://www.zhenai.com/zhenghun/alashanmeng
	//City: 阿勒泰, URL: http://www.zhenai.com/zhenghun/aletai
	//City: 阿里, URL: http://www.zhenai.com/zhenghun/ali
	//City: 安徽, URL: http://www.zhenai.com/zhenghun/anhui
	// ...

}

// IDEA函数返回值自动补全 option+command+v

// 元素过滤
	// CSS选择器
	// xpath 能力没有CSS选择器强
	// 正则表达式