//通过共享变量上锁来构建并发结构
package memo4

import "sync"

//Func是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//对key的第一次访问，这个goroutine负责计算数据和广播数据
		//已准备完毕的消息
		e = &entry{ready:make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		//广播数据已准备完毕的消息
		close(e.ready)
	} else {
		//对这个key的重复访问
		memo.mu.Unlock()
		//等待数据准备完毕（参考8.9取消操作，利用通道关闭来广播一个事件）
		<- e.ready
	}

	return e.res.value, e.res.err
}