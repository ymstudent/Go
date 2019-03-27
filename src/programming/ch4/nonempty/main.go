package main

import "fmt"

func main()  {
	data := []string{"a", "", "b"}
	/*res1 := nonempty(data)
	fmt.Println(res1)*/
	res2 := nonempty2(data)
	fmt.Println(res2)
}

func nonempty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[0:i]
}

func nonempty2(strings []string) []string {
	for i, v := range strings {
		if v != "" {
			fmt.Println(strings)
			strings[i] = v
		}
	}
	return strings[:]
}
