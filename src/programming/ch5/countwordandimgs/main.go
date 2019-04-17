//统计页面字数与图片
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main()  {
	words, imgs, err := CountWordsAndImgs("https://ymfeb.cn")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(words, imgs)
}

func CountWordsAndImgs(url string) (words, imgs int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s\n", err)
		return
	}
	words, imgs = count(doc)
	return
}

func count(n *html.Node) (words, imgs int) {
	if n.Type == html.TextNode {
		scaner := bufio.NewScanner(strings.NewReader(n.Data))
		scaner.Split(bufio.ScanWords)
		for scaner.Scan() {
			words++
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		imgs++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ws, is := count(c)
		words += ws
		imgs += is
	}
	return
}
