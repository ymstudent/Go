package main

import (
	"fmt"
	"sync"
)

//使用Go的并发特性来并行优化一个Web爬虫
//提示；可以用一个map来缓存已经获取的URL，但需要注意map本身并不是并发安全的

type Fetcher interface {
	//Fetch返回URL的body内容，并且将这个页面内找到的URL放入一个slice中
	Fetch(url string) (body string, urls []string, err error)
}

//使用map缓存已经获取的URL
type SafeMap struct {
	v map[string]int
	mux sync.Mutex
}

func Crawl(url string, depth int, fetcher Fetcher)  {
	//TODO：并行的抓取URL
	//TODO：不重复抓取页面
	if depth <= 0 {
		return
	}

	//避免重复抓取
	if checkUrl(url) {
		return
	}

	//标记已抓取页面
	maskUrlCount(url, 1)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	//并行抓取
	ch := make(chan int)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			ch <- 1
		}(u)
	}
	//等待抓取完成
	for range urls {
		<- ch
	}

	return
}

// UrlCounter 的并发使用是安全的。
type UrlCounter struct {
	v   map[string]int
	mux sync.Mutex
}


var urlCounter UrlCounter

func checkUrl(url string) bool {
	urlCounter.mux.Lock()
	defer urlCounter.mux.Unlock()
	_, ok := urlCounter.v[url]
	// if ok && c == 0 {
	// 	return false
	// }
	return ok
}

func maskUrlCount(url string, count int) {
	urlCounter.mux.Lock()
	defer urlCounter.mux.Unlock()
	urlCounter.v[url] += count
}

func main()  {
	urlCounter = UrlCounter{v: make(map[string]int)}
	Crawl("https://golang.org/", 4, fetcher)
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	//判断映射中是否存在对应url
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string {
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}


