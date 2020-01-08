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
	"strings"
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

func (spider Spider) GetGithub() []map[string]interface{} {
	var allData []map[string]interface{}

	url := "https://github.com/trending"
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

	request.Header.Add("cookie", `_octo=GH1.1.1197201112.1563518143; _ga=GA1.2.2103263047.1563518145; _device_id=bdd17a569cbadf10478c9999c3673f4c; user_session=_bZNm-H_tP3XtvBkaN5hTD-KP2QbvkWJkmI_hFZVl4KXuQth; __Host-user_session_same_site=_bZNm-H_tP3XtvBkaN5hTD-KP2QbvkWJkmI_hFZVl4KXuQth; logged_in=yes; dotcom_user=ymstudent; tz=Asia%2FShanghai; has_recent_activity=1; _gh_sess=TFkvTEpZekNnUitZSzdzem03VS9HLzQ4RGZVMkdETmI4WTJEU0QxTXAzOFdTcFhxNjgwOGRLYlhaZndGeDVjSGlEWjVLMEdCZjFSRlI4MHdmSzZqdEpiRG16a2dFRHdad0RKZldZN1lwN1ZKWnh5Qm96UXlZRWdIOHU0ckZ2MUxEdzdzTXFNM0lSNlRGWlV1OEhYSDNsSjNUcWx6TWVJWnhHb3MyTkdqak1hdS9yVWVWakYzNldCZmU5U0wxVWdWeGpsWlFaZElGSjJlNXdjYStuelMwekh2dUVuMy9ucWtTc21Lcy9iYlYwYzRZMmVQSHVYVjdxVGxFR1E2YmtrUVlUWDB4WHRIdjVqUURxd2ZvVlJTc1UzcmVxTmw0QTBMWS9XbnFjRENBaVJyY1RkbkZPMzFwTGtRRlBReGR5eXFzWVUxbWd2ajdMT0NQU0Ewb2NXZTdyMDNwTS8zcTJDOG1ZZ2ZlNnBTMWhaTzh0ZnluSWhrNG9TTHA0WWp3MFk4dTA4V2Q4OUhpZjZpbnVYeFdlQnRtZ21CZ3dBcFo1Qm9oTnczc2RQT3NxMkR5KzZTU1REeThxU1BBMDZJRHkrdStoOGM2b0ZIL09yOENZclNaSDNFSU1ScE5uWXZ4RVZibXpWZ3dWUUlKRkNtQmxDNEVacXN6UUpkdk04U05TZVhJdnlTYW56cDY4aGJ0eVlJczJ2U2t0TDNBb25HMnVhQ1ova2JGVm41VFBReENVNlRkQWgrUFNzZUZhQ1FIUmhWZ1AyWXZneTFodG5aeEhhZmsyUzdWY05rREdCbHJxRDlCL1Fyd2ZWU2dnQUo1Zm9xMlhnRG5YaFQ0TytxZUszRGFZbzhOWUdzWEpTbEgyVE05dGtMVWFBVFFDWjliQ1l5MmxiaStMRkdyT0RyMFgzckRXV2VZTmhVWFNMTEk5RFhBSmpxdVBtM1pXR3pJcmdCdDMvRHhzZ01hb3BDcUJyRjkzalhWUHFtb2dabU1ScUNmaXhzU21iQU02dE1BZEJETE5HUkRRbkt4TFBRcXBxeWlYV2RvWHlyekhST0JtMlpBclVKQ1JyblU5T1lUQ2c2QU1QU3VLVU44TEk5UzJ1T2Z4SzhnNDBVdUlmRjQrVU1ObWIxRHgrcEtQczVPT0J1RHE4TTFZME5MV0c2bW45dGhpcEpub1I4ckhKVGZCVmo2NGFlRFZkVTVER2RnVHlZTlRnR2g3Q09yeG43OEZRMXRmN2N2SlNPYmtCOUZOZVJwMndQOUFyM0VMOGo2TUNvYXlzWWNLZTluWkRFRERLR2pueGhUQWhoR2xmR3dLb3lrZ2p4T3lVWWllVmVDNXliOWtXQlo1UzVFc1d0a3diSGhna2JHL3IrSU95bUUrVUJGbXZ0Z2pCNzE3RExRL1pXUTAwMDB5azczMytOS1hOdzVhL0xhdFg4bTMvaGNxaW45MDVpVk5kbUxIb1RjZUFwcm9zb2ZycGY2MEpESFBlM0VRT3d3RGJ4YjR2d2FrQ0pucldvM3dEZHFFRXo4Wk5keWpNUWk5TWgvYlNOdGhtaFVmc2FEZzV3SkxvckFsTHErenRyc2FXakFDdz0tLWZac2VnVEJheUJwQVhCY203SCtVd1E9PQ%3D%3D--7f385eada410f979251421b3609ff8f036c2abdd`)
	request.Header.Add("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36`)
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

	document.Find(".Box .Box-row").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find(".text-gary").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": strings.Trim(text, ""), "url":"https://github.com/"+url})
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
	spider := Spider{DataType: "Github"}
	start := time.Now()
	group.Add(1)
	go ExecGetData(spider)
	group.Wait()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费%.2fs秒完成抓取%s", seconds, spider.DataType)
	fmt.Println("抓取完成")
}
