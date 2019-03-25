//一个并发请求程序
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main()  {
	url := "http://localhost:8000"
	for i :=0; i < 2; i++ {
		go func(url *string) {
			_, err := http.Get(*url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
		}(&url)
	}
}
