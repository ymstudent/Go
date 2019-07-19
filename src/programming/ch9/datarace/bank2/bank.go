package bank2
//避免数据竞态的方法3：保证同一时间最多只有一个goroutine能访问共享变量
//使用二进制信号量保证同一时间最多只有一个goroutine能访问共享变量
var (
	sema = make(chan struct{}, 1)	//用来保护balance的二进制信号量
	balance int
)

func Deposit(amount int)  {
	sema <- struct{}{}	//获取令牌
	balance = balance + amount
	<-sema	//释放令牌
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<- sema
	return b
}
