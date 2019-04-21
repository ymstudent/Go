//defer语句：将函数或方法推迟到函数执行完毕前或返回错误前执行，保证资源能正常释放
//一般用于成对操作，如打开/关闭，连接/断开，加锁/解锁。通常在成功获得资源时调用
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"programming/ch5/links"
	"strings"
)

func main()  {

}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //使用defer语句关闭资源
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild)
		}
	}

	links.ForeachNode(doc, visitNode, nil)
	return nil
}
