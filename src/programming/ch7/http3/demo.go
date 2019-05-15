//请求多工转发器serverMux
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

type database map[string]dollars

func main()  {
	db := database{"shoes":50, "socks":5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list)) //注册多个http.Handle
	mux.Handle("/price",http.HandlerFunc(db.price))
	mux.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w,"%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w,"%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "invalid price")
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "update sucess")
}


