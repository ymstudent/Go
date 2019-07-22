package memo1_test

import (
	"fmt"
	"gowork/programming/ch9/memo1"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string  {
	URLs := []string{"https://www.ymfeb.cn","https://www.baidu.com","https://400.semoor.cn",
		"https://www.ymfeb.cn","https://www.baidu.com","https://400.semoor.cn"}
	return  URLs
}

func TestMemo(t *testing.T)  {
	m := memo1.New(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}

}
