package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie {
	{
		Title:"Casablanca", Year:1942, Color:false, Actors:[]string {"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:"Cool Hand Luke", Year:1967, Color:true, Actors:[]string {"Paul Newman"},
	},
	//.....
}

var titles []struct{Title string}

func main()  {
	//data, err := json.Marshal(movies) //将GO数据结构转换为JSON字符串
	data, err := json.MarshalIndent(movies, "", "	") //将GO数据结构转换为格式化后的JSON字符串
	if err != nil {
		fmt.Printf("Json marshal failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)

	if err := json.Unmarshal(data, &titles); err != nil { //将JSON字符串解码到结构体slice titles中
		fmt.Printf("Json unmarshal failed: %s\n", err)
	}
	fmt.Println(titles)
}
