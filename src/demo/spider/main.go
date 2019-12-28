package main

/**
 * 爬虫代码：来自今日热榜：https://github.com/tophubs/TopList
 */
import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"reflect"
	"sync"
	"time"
)

type HotData struct {
	Code    int
	Message string
	Data    interface{}
}

type Spider struct {
	DataType string
}

func (spider Spider) GetV2EX() []map[string]interface{} {
	url := "https://www.v2ex.com/?tab=hot"
	timeOut := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeOut,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}

	var allData []map[string]interface{}
	document.Find(".item_title").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.v2ex.com" + url})
		}
	})
	return allData
}

func SaveDataToJson(data interface{}) string {
	Message := HotData{}
	Message.Code = 0
	Message.Message = "获取成功"
	Message.Data = data
	jsonStr, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列化JSON错误")
	}
	return string(jsonStr)
}

func ExecGetData(spider Spider) {
	defer group.Done()
	reflectValue := reflect.ValueOf(spider)
	dataType := reflectValue.MethodByName("Get" + spider.DataType)
	data := dataType.Call(nil)
	originData := data[0].Interface().([]map[string]interface{})
	fmt.Println(SaveDataToJson(originData))
}

var group sync.WaitGroup

func main() {
	spider := Spider{DataType: "V2EX"}
	start := time.Now()
	group.Add(1)
	go ExecGetData(spider)
	group.Wait()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费%.2fs秒完成抓取%s", seconds, spider.DataType)
	fmt.Println("抓取完成")
}
