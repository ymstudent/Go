package eval

import (
	"fmt"
	"math"
)

//表示一个变量，如x
type Var string

//表示一个数字常量，如3.14
type literal float64

//表示一元操作符表达式，如-x
type unary struct {
	op rune //+,-中的一个
	x  Expr
}

//表示二元操作符表达式，如x+y
type binary struct {
	op   rune //+,-,*,/中的一个
	x, y Expr
}

//表示函数调用表达式，如sin(x)
type call struct {
	fn   string //pow, sin, sqrt中的一个
	args []Expr
}

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}
