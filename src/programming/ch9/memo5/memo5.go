//使用通信顺序进程构建并发结构
package memo5

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

//request 是一条请求信息， key 需要用 Func 来调用
type request struct {
	key string
	response chan <- result //客户端需要单个result
}

type Memo struct {
	requests chan request
}

//New返回f的函数记忆，客户端之后需要调用close
func New(f Func) *Memo {
	memo := &Memo{requests:make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready:make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	//执行函数
	e.res.value, e.res.err = f(key)
	//同志数据已经准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan <- result) {
	//等待该数据准备完毕
	<- e.ready
	//向客户端发送结果
	response <- e.res
}
