package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	rand2 "math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	filePath := flag.String("f", "", "文件路径")
	tplId := flag.String("t", "", "模版ID")
	flag.Parse()

	start := time.Now()
	csvFile, err := os.Open(*filePath)
	check(err)
	defer csvFile.Close()

	var wg sync.WaitGroup
	csvReader := csv.NewReader(csvFile)
	arr, _ := csvReader.ReadAll()
	counter := make(chan bool, len(arr)+1)
	limit := make(chan bool, 1000) //用来限制goroutine数量
	pk, err := ParsePrivateKey()
	check(err)

	for _, row := range arr {
		wg.Add(1)
		limit <- true
		go func(row []string) { //通过添加显式参数，确保当go语句执行时，使用当前row值（参考5.6.1内部匿名函数中获取循环变量的问题）
			defer wg.Done()
			success, err := sendMsg(row, *tplId, pk)
			if err != nil {
				fmt.Println(err)
			}
			counter <- success
			<-limit
		}(row)
	}

	go func() {
		wg.Wait()
		close(counter) //安全关闭通道
	}()

	//统计成功与失败数量
	successNum := 0
	failNum := 0
	for v := range counter {
		if v {
			successNum++
		} else {
			failNum++
		}
	}
	fmt.Printf("发券成功:%d\n", successNum)
	fmt.Printf("发券失败:%d\n", failNum)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func sendMsg(row []string, tplId string, pk *rsa.PrivateKey) (success bool, err error) {
	query := getQuery(row, tplId, pk)
	queryUrl := "https://openapi.alipay.com/gateway.do?" + query
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //跳过https证书验证
	}
	c := &http.Client{
		Transport: tr,
	}
	resp, err := c.Get(queryUrl)
	if err != nil {
		err := fmt.Errorf("请求错误:%s", err)
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("请求失败,错误码:%d", resp.StatusCode)
		return false, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("读取返回结果失败:%s", err)
		return false, err
	}
	var smsresp map[string]interface{}
	if err := json.Unmarshal(body, &smsresp); err != nil {
		err := fmt.Errorf("JSON解析失败:%s", err)
		return false, err
	}
	res := smsresp["alipay_pass_instance_add_response"].(map[string]interface{})
	if res["code"] == "10000" {
		return true, nil
	} else {
		return false, nil
	}
}

func getQuery(row []string, tplId string, pk *rsa.PrivateKey) string {
	recognitionInfoMap := make(map[string]string)
	recognitionInfoMap["partner_id"] = strings.Trim(row[0], "`")
	recognitionInfoMap["out_trade_no"] = strings.Trim(row[1], "`")
	recognitionInfo, _ := json.Marshal(recognitionInfoMap)

	tplParamsMap := make(map[string]int)
	tplParamsMap["channelID"] = 123456
	r := rand2.New(rand2.NewSource(time.Now().Unix()))
	tplParamsMap["serialNumber"] = r.Intn(100)
	tplParams, _ := json.Marshal(tplParamsMap)

	body := make(map[string]string)
	body["tpl_id"] = tplId
	body["recognition_type"] = "1"
	body["tpl_params"] = string(tplParams)
	body["recognition_info"] = string(recognitionInfo)

	data := url.Values{}
	bizContent, _ := json.Marshal(body)
	data.Set("app_id", "123456")
	data.Set("biz_content", string(bizContent))
	data.Set("charset", "utf-8")
	data.Set("format", "JSON")
	data.Set("method", "alipay.pass.instance.add")
	data.Set("sign_type", "RSA2")
	data.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	data.Set("version", "1.0")
	signContentBytes, _ := url.QueryUnescape(data.Encode())
	sign := sign([]byte(signContentBytes), pk)
	data.Set("sign", sign)
	return data.Encode()
}

func sign(data []byte, pk *rsa.PrivateKey) string {
	h := sha256.New()
	hType := crypto.SHA256
	h.Write(data)
	d := h.Sum(nil)
	bs, err := rsa.SignPKCS1v15(rand.Reader, pk, hType, d)
	check(err)
	return base64.StdEncoding.EncodeToString(bs)
}

func ParsePrivateKey() (pk *rsa.PrivateKey, err error) {
	privateKey := ``
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = fmt.Errorf("私钥格式错误1:%s", privateKey)
		return
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err == nil {
			pk = rsaPrivateKey
		} else {
			err = fmt.Errorf("私钥格式错误2:%s", privateKey)
		}
	default:
		err = fmt.Errorf("私钥格式错误:%s", privateKey)
	}
	return
}


