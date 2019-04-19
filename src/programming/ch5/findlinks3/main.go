package main

import (
	"fmt"
	"log"
	"programming/ch5/links"
)

func main()  {
	worklist := []string{"http://ymfeb.cn"}
	breadthFirst(crawl, worklist)
}

func breadthFirst(f func(item string) []string, worklist []string)  {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
