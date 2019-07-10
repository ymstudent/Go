package main

import (
	"fmt"
	"programming/ch9/datarace/bank"
	"time"
)

//数据竞态发生于多个goroutine并发读写同一个变量并且至少其中一个是写入时
func main()  {
	for i := 0; i < 10; i++ {
		bank.ReturnZero()
		dataRaceTest()
		time.Sleep(time.Second)
	}
}

//正常情况下输出= 300，但如果发生数据竞态（data race），输出为= 200，Bob存入的钱消失了
func dataRaceTest()  {
	//Alice
	go func() {
		bank.Deposit(200) 					//A1
		fmt.Println("=", bank.Balance())		//A2
	}()
	//Bob
	go bank.Deposit(100)						//B
}
