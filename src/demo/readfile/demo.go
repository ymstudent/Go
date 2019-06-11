package main

import (
	"bufio"
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
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	counter := make(chan int)
	csvFile, err := os.Open("./dm.csv")
	check(err)

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		go sendMsg(row, counter)
	}
}

func sendMsg(row []string, counter chan int)  {
	query := getQuery(row)
	queryUrl := "https://openapi.alipay.com/gateway.do?"+query
	fmt.Println(queryUrl)
	resp, err := http.Get(queryUrl)
	check(err)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("请求失败,错误码:%d", resp.StatusCode))
	}
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getQuery(row []string) string {
	aa := make(map[string]string)
	aa["partner_id"] = strings.Trim(row[0], "`")
	aa["out_trade_no"] = strings.Trim(row[1], "`")

	bb := make(map[string]string)
	bb["channelID"] = "123456"
	bb["serialNumber"] = md5V(aa["out_trade_no"])

	body := make(map[string]string)
	body["tpl_id"] = "123456"
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
	data.Set("app_id", "123456")
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


