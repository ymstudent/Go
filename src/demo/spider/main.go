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

//抓取知乎
func (spider Spider) GetZhiHu() []map[string]interface{} {
	url := "https://www.zhihu.com/hot"
	timeOut := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout:timeOut,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}

	request.Header.Add("Cookie", `_zap=33838d24-4974-4107-9a3e-9bb5684d9b60; d_c0="ANAiE4NRwg-PTsk0jHaOMj4SoKyKtByS-ew=|1563518525"; _xsrf=ZBvxuEBek4cBqJ2W5LSFuSK56MNe1Bnn; q_c1=ae80f4927e694370a3fd48939e56fcb2|1574000335000|1574000335000; __utma=51854390.560013395.1574606799.1574606799.1574606799.1; __utmz=51854390.1574606799.1.1.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmv=51854390.100-1|2=registration_date=20181103=1^3=entry_date=20181103=1; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1577579430,1578215032,1578216564,1578323899; capsion_ticket="2|1:0|10:1578323923|14:capsion_ticket|44:ZTRkNmYyNmVmZTAyNGVlN2E0YTY5ZTk2NmQ5OGM0NTk=|ba1a5364ffb434a2b4459c99ae6315c9fc3f3fcd43d5344d2137292e7342d345"; z_c0="2|1:0|10:1578323949|4:z_c0|92:Mi4xVzFVSERRQUFBQUFBMENJVGcxSENEeWNBQUFDRUFsVk43ZHc2WGdDR2VHUkZORmswemZfT2RkcFhQU1VMcGs5VnRB|a4cc29c0decf0d99701fc3ced7fac4a19ded0248542663600707ae6dd26340f4"; tshl=; tst=h; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1578324049; KLBRSID=9d75f80756f65c61b0a50d80b4ca9b13|1578325500|1578323898`)
	request.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36`)
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
	document.Find(".HotList-list .HotItem-content").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("h2").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.zhihu.com" + url})
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
	spider := Spider{DataType: "ZhiHu"}
	start := time.Now()
	group.Add(1)
	go ExecGetData(spider)
	group.Wait()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费%.2fs秒完成抓取%s", seconds, spider.DataType)
	fmt.Println("抓取完成")
}
