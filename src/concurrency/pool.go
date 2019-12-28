package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("cannot dial host: %v", err)
	}
	if _, err := ioutil.ReadAll(conn); err != nil {
		log.Fatalf("cannot read: %v", err)
	}
	conn.Close()
}

//模拟创建一个连接到服务器的函数
func connectToServer() interface{} {
	time.Sleep(1*time.Second)
	return struct {}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("cannot accept connection: %v", err)
				continue
			}
			fmt.Println("sleep 1s")
			connectToServer()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}


func warmServerConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToServer,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon2() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServerConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}





