package bank1
//如何避免发生数据竞态2：避免从多个goroutine访问同一个变量
var deposits = make(chan int) 	//发送存款额
var balances = make(chan int)	//接收余额

func Deposit(amount int)  {
	deposits <- amount
}

func Balance() int {
	return <- balances
}

func teller()  {
	var balance int		//balance被限制在teller goroutine中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init()  {
	go teller()
}
