package main

import "fmt"

//import "golang.org/x/tour/reader"

//练习Reader:实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。

type MyReader struct {}

func (r MyReader) Read(b []byte) (n int, err error)  {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main()  {
	//reader.Validate(MyReader{})
	r := MyReader{}
	b := make([]byte, 4)
	n, err := r.Read(b)
	fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
}
