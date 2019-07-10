package bank

var balance int

func ReturnZero()  {
	balance = 0
}

func Deposit(amount int)  {
	balance = balance + amount  //此处实际是两个串行操作，可以分为读部分（balance = ）和写部分（= balance + amount）
}

func Balance() int {
	return balance
}