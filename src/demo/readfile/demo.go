package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
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
	start := time.Now()
	csvFile, err := os.Open("./dm.csv")
	check(err)
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	arr, _ := csvReader.ReadAll()
	counter := make(chan string)
	var wg sync.WaitGroup //工作goroutine的个数
	for _, row := range arr {
		wg.Add(1)
		go func(row []string){  //通过添加显式参数，确保当go语句执行时，使用当前row值（参考5.6.1内部匿名函数中获取循环变量的问题）
			defer wg.Done()
			params := getQuery(row)
			counter <- params
		}(row)
	}
	go func() {
		wg.Wait()
		close(counter) //安全关闭通道
	}()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	var s []string
	for v := range counter {
		s = append(s, v)
	}
	fmt.Println(len(s))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func sendMsg(row []string) (success bool, err error) {
	query := getQuery(row)
	queryUrl := "https://openapi.alipay.com/gateway.do?"+query
	resp, err := http.Get(queryUrl)
	if err != nil {
		err := fmt.Errorf("请求失败:%s", err)
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
		err := fmt.Errorf("JSON解析失败:%s", res)
		return false, err
	}
}

func getQuery(row []string) string {
	aa := make(map[string]string)
	aa["partner_id"] = strings.Trim(row[0], "`")
	aa["out_trade_no"] = strings.Trim(row[1], "`")

	bb := make(map[string]string)
	bb["channelID"] = ""
	bb["serialNumber"] = md5V(aa["out_trade_no"])

	body := make(map[string]string)
	body["tpl_id"] = ""
	tpl_params, err := json.Marshal(bb)
	check(err)
	body["tpl_params"] = string(tpl_params)
	body["recognition_type"] = "1"
	recognition_info, err := json.Marshal(aa)
	check(err)
	body["recognition_info"] = string(recognition_info)

	data := url.Values{}
	biz_content, err := json.Marshal(body)
	check(err)
	data.Set("app_id", "")
	data.Set("biz_content", string(biz_content))
	data.Set("charset", "utf-8")
	data.Set("format", "JSON")
	data.Set("method", "alipay.pass.instance.add")
	data.Set("sign_type", "RSA2")
	data.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	data.Set("version", "1.0")
	signContentBytes, _ := url.QueryUnescape(data.Encode())
	sign := sign([]byte(signContentBytes))
	data.Set("sign", sign)
	return data.Encode()
}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sign(data []byte) string {
	prikey := ``


	h := sha256.New()
	hType := crypto.SHA256
	h.Write(data)
	d := h.Sum(nil)
	pk, err := ParsePrivateKey(prikey)
	check(err)
	bs, err := rsa.SignPKCS1v15(rand.Reader, pk, hType, d)
	check(err)
	return base64.StdEncoding.EncodeToString(bs)
}


func ParsePrivateKey(privateKey string) (pk *rsa.PrivateKey, err error) {
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


