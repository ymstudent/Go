package main

import (
	"fmt"
	"programming/ch9/datarace/bank"
	"programming/ch9/datarace/bank1"
	"time"
)

//数据竞态发生于多个goroutine并发读写同一个变量并且至少其中一个是写入时
func main()  {
	for  {
		bank.ReturnZero()
		dataRaceTest()
		time.Sleep(100 * time.Millisecond)
		if  bank.Balance() == 200 {
			fmt.Println("occur data race")
			break
		}
	}
	time.Sleep(10 * time.Second)
}

func noDataRaceTest() {
	go func() {
		bank1.Deposit(200)
		fmt.Println("=", bank1.Balance())
	}()

	go bank1.Deposit(100)
}

//正常情况下输出= 300，但如果发生数据竞态（data race），输出为= 200，Bob存入的钱消失了
func dataRaceTest()  {
	//Alice
	go func() {
		bank.Deposit(200) 					//A1
		//fmt.Println("=", bank.Balance())		//A2
	}()
	//Bob
	go bank.Deposit(100)						//B
}
