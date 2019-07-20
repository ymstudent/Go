package bank3

//使用互斥锁来避免数据竞态
import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int)  {
	mu.Lock()			//加锁
	defer mu.Unlock()	//释放锁
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}


//错误的实现：无法对一个已经上锁的互斥量再上锁。这里会导致死锁
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false	//余额不足
	}
	return true
}

//优化：将Deposit函数拆分：一个不导出的deposit，以及一个导出的Deposit2(专门用来执行上锁，解锁操作)。
func Deposit2(amount int)  {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func deposit(amount int)  {
	balance += amount
}

func Withdraw2(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

