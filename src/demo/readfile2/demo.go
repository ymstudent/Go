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
	"io"
	"io/ioutil"
	rand2 "math/rand"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//解析参数
	filePath := flag.String("f", "", "文件路径")
	tplId := flag.String("t", "", "模版ID")
	flag.Parse()

	//解析密钥
	pk, err := ParsePrivateKey()
	check(err)

	//读取文件
	start := time.Now()
	paramsChan := make(chan string, runtime.NumCPU())
	go readFile(*filePath, *tplId, pk, paramsChan)

	//发送数据
	var failNum int64
	var successNum int64
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for s := range paramsChan {
				res, err := sendMsg(s)
				if res {
					atomic.AddInt64(&successNum, 1)
				} else {
					if err != nil {
						fmt.Println(err)
					}
					atomic.AddInt64(&failNum, 1)
				}
			}
		}()
	}
    wg.Wait()

	fmt.Printf("发券成功:%d\n", successNum)
	fmt.Printf("发券失败:%d\n", failNum)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readFile(filePath string, tplId string, pk *rsa.PrivateKey, paramsChan chan string) {
	csvFile, err := os.Open(filePath)
	check(err)
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	limit := make(chan struct{}, runtime.NumCPU()) //计数信号量
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("读取csv错误: %s\n", err)
		}
		limit <- struct{}{}
		go func() {
			defer func() {
				<-limit
			}()
			params, err := getQuery(row, tplId, pk)
			if err != nil {
				fmt.Println(err)
			}
			paramsChan <- params
		}()
	}
	for i := 0; i < cap(limit); i++ {
		limit <- struct{}{}
	}
	close(paramsChan)
}

func sendMsg(query string) (success bool, err error) {
	queryUrl := "https://openapi.alipay.com/gateway.do?" + query
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //跳过https证书验证
	}
	c := &http.Client{
		Transport: tr,
	}
	resp, err := c.Get(queryUrl)
	if err != nil {
		err = fmt.Errorf("请求错误: %s", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("请求失败,错误码: %d", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("读取返回结果失败: %s", err)
		return
	}
	var smsresp map[string]interface{}
	if err = json.Unmarshal(body, &smsresp); err != nil {
		err = fmt.Errorf("JSON解析失败: %s", err)
		return
	}
	res := smsresp["alipay_pass_instance_add_response"].(map[string]interface{})
	if res["code"] == "10000" {
		return true, nil
	} else {
		return false, nil
	}
}

func getQuery(row []string, tplId string, pk *rsa.PrivateKey) (query string, err error) {
	recognitionInfoMap := make(map[string]string)
	recognitionInfoMap["partner_id"] = strings.Trim(row[0], "`")
	recognitionInfoMap["out_trade_no"] = strings.Trim(row[1], "`")
	recognitionInfo, err := json.Marshal(recognitionInfoMap)
	if err != nil {
		err = fmt.Errorf("recognitionInfoMap JSON加密失败: %s", err)
		return
	}

	tplParamsMap := make(map[string]int)
	tplParamsMap["channelID"] = 123456
	r := rand2.New(rand2.NewSource(time.Now().Unix()))
	tplParamsMap["serialNumber"] = r.Intn(100)
	tplParams, err := json.Marshal(tplParamsMap)
	if err != nil {
		err = fmt.Errorf("tplParamsMap JSON加密失败: %s", err)
		return
	}

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

	signContentBytes, err := url.QueryUnescape(data.Encode())
	if err != nil {
		err = fmt.Errorf("url QueryUnescape失败: %s", err)
		return
	}

	sign, err := sign([]byte(signContentBytes), pk)
	if err != nil {
		err = fmt.Errorf("签名加密失败: %s", err)
		return
	}

	data.Set("sign", sign)
	query = data.Encode()
	return
}

func sign(data []byte, pk *rsa.PrivateKey) (sign string, err error) {
	h := sha256.New()
	hType := crypto.SHA256
	h.Write(data)
	d := h.Sum(nil)
	bs, err := rsa.SignPKCS1v15(rand.Reader, pk, hType, d)
	if err != nil {
		err = fmt.Errorf("rsa SignPKCS1v15失败: %s", err)
		return
	}
	sign = base64.StdEncoding.EncodeToString(bs)
	return
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