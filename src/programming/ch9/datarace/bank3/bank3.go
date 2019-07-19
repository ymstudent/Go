package bank3

//使用互斥锁来避免数据竞态
import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int)  {
	mu.Lock()	//加锁
	defer mu.Unlock()	//释放锁
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
