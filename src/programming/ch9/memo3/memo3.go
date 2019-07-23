package memo3

import "sync"

//Func是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

//非阻塞缓存，将获取与更新缓存操作分别加锁，但在进行Get请求的地方是可以并行的
//缺点：多个goroutine调用同一个URL时，会同时更新缓存，其中一个结果会被另一个覆盖（等于多了一次额外处理）
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key) //此处多个goroutine可并行
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}