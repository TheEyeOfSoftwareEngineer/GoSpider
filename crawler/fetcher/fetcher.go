package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
/**
	输入: URL
	输出: 对应文档的字节流或者错误
 */
func Fetch(url string) ([]byte, error) {

	time.Sleep(time.Duration(2)*time.Second)

	client := &http.Client{}
	//response, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	cookie := "Your Cookie"
	request.Header.Add("cookie", cookie)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusAccepted {
		time.Sleep(time.Duration(5)*time.Second)
		fmt.Println("继续访问")
		Fetch(url)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
	}

	bodyReader := bufio.NewReader(response.Body)
	e := determineEncoding(bodyReader)
	//encoding转换
	newReader := transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(newReader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	//猜测获取HTML的encoding
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}