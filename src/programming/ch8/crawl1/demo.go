package main

import (
	"log"
	"os"
	"programming/ch5/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main()  {
	worklist := make(chan []string)
	var n int //等待发送到任务列表的数量

	//从命令行参数开始
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
