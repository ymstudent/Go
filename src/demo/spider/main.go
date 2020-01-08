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

func (spider Spider) GetHuPu() []map[string]interface{} {
	var allData []map[string]interface{}
	url := "https://bbs.hupu.com/all-gambia"
	timeOut := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout:timeOut,
	}

	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return allData
	}

	request.Header.Add("cookie", `_dacevid3=2762cd3f.a0e0.2d82.7f4b.465bf2d87b2b; __gads=ID=8ea0353ce6a5e604:T=1564304070:S=ALNI_MbyxkLdavVvYxxq_bGgYFS8E1OPmw; _HUPUSSOID=6888bdc7-0cbe-4bfb-ab42-dcf33bb00909; acw_tc=76b20f6015762909611003634e49681626e0f21d69f7bc6bed60300ec745e7; _CLT=b0c2a05996d8b48b354e1fa4ddfc1fef; u=59791349|6JmO5omRSlIwOTE2OTAzODE1|c949|fbc38e6f125b5c142fa3d7a89f67d053|125b5c142fa3d7a8|aHVwdV9iNzA2Mjk5NWQ2YmVjNGFj; us=d00a4e8d59d63ee520b7811cc17799ab8b78d6a6ad246fbe09a3333f1503e9e6738764ac52d7953cf68f3c8e5cdb163128aaddd79dab6671868c59ec808b72fc; ua=31569934; PHPSESSID=6a5103eb0828848b27da7e041c3fd844; _cnzz_CV30020080=buzi_cookie%7C2762cd3f.a0e0.2d82.7f4b.465bf2d87b2b%7C-1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2216c90b892e2bf-0bf641a103ead5-38637701-1440000-16c90b892e3df%22%2C%22%24device_id%22%3A%2216c90b892e2bf-0bf641a103ead5-38637701-1440000-16c90b892e3df%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E5%BC%95%E8%8D%90%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Fmo.fish%2Fmain%2Fhome%2Fhot%22%2C%22%24latest_referrer_host%22%3A%22mo.fish%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%7D%7D; UM_distinctid=16f85be09d542e-03979c831bc09a-1d336b5a-13c680-16f85be09d68bb; Hm_lvt_3d37bd93521c56eba4cc11e2353632e2=1578496822; Hm_lpvt_3d37bd93521c56eba4cc11e2353632e2=1578496822; Hm_lvt_39fc58a7ab8a311f2f6ca4dc1222a96e=1577021178,1578496693,1578496805,1578496833; _fmdata=2tzKYiaN7nEtlq%2FPE90sOpu95aWJPWUeIs9%2B%2F9Yy9jdPZiOiHsPZcQn%2BgatZtEXt4CQMLfODXxgq38VJrcV%2BZs9H78wBJCjSsiAhPXO0vyQ%3D; __dacevst=ad8acee6.80575057|1578498919081; Hm_lpvt_39fc58a7ab8a311f2f6ca4dc1222a96e=1578497119`)
	request.Header.Add("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return allData
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return allData
	}

	document.Find(".bbsHotPit .list").First().Find(".textSpan").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text, _ := selection.Find("a").Attr("title")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title":text, "url":"https://bbs.hupu.com"+url})
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
	spider := Spider{DataType: "HuPu"}
	start := time.Now()
	group.Add(1)
	go ExecGetData(spider)
	group.Wait()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费%.2fs秒完成抓取%s", seconds, spider.DataType)
	fmt.Println("抓取完成")
}
