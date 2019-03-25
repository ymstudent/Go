package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
//映射的文法与结构体相似，不过必须有键名。
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google": Vertex{37.42202, -122.08408},
}
func main()  {
	fmt.Println(m)
	fmt.Println(m2)
}