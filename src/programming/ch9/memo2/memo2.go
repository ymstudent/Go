package memo2

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

//并发安全，但是每次调用f都加锁，导致Get操作其实是串行的，性能低
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
