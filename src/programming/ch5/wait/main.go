//五种常用错误处理策略
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main()  {
	err := WaitForServer("https://golang.org")
	if err != nil {
		//3、输出错误并优雅的停止程序
		log.SetPrefix("wait: ")  //自己定义命令的名称作为log包的前缀
		//log.SetFlags(0) //将日期与时间略去
		log.Fatalf("server is down: %v\n", err) //默认将时间与日期作为前缀添加到错误消息前
		os.Exit(1)
	}

	//4、记录下错误信息并继续运行程序
	log.Printf("wait: %v", err)
	//5、在某些特殊的情况下可以直接安全的忽略掉整个日志
	os.RemoveAll("../../ch5") //忽略错误
}

//尝试连接url对应的服务器，在一分钟内使用指数退避策略进行重试，所有尝试失败后返回错误
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	//2、对不可预测的错误，在短暂的间隔后对操作进行重试，超出一定重试次数与时间后再报错退出
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //成功
		}
		log.Printf("server not responding (%s); retrying.....", err)//1、为原始错误消息添加上下文信息来建立一个可读的错误描述
		//fmt.Println(time.Second << uint(tries)) 1, 2, 4.....
		time.Sleep(time.Second << uint(tries)) //指数退避策略
	}
	return fmt.Errorf("server %s faild to respond after %s", url, timeout)
}
