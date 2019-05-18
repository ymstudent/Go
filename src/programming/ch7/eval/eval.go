package eval

import (
	"fmt"
	"math"
	"testing"
)

//表示一个变量，如x
type Var string
//表示一个数字常量，如3.14
type literal float64
//表示一元操作符表达式，如-x
type unary struct {
	op rune //+,-中的一个
	x Expr
}
//表示二元操作符表达式，如x+y
type binary struct {
	op rune //+,-,*,/中的一个
	x, y Expr
}
//表示函数调用表达式，如sin(x)
type call struct {
	fn string //pow, sin, sqrt中的一个
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
		return b.x.Eval(env)+b.y.Eval(env)
	case '-':
		return b.x.Eval(env)-b.y.Eval(env)
	case '*':
		return b.x.Eval(env)*b.y.Eval(env)
	case '/':
		return b.x.Eval(env)/b.y.Eval(env)
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

func TestEval(t *testing.T)  {
	tests := []struct{
		expr string
		env Env
		want string
	}{
		{"sqrt(A/pi)", Env{"A":87616, "pi":math.Pi}, "167"},
		{"pow(x, 3)+pow(y, 3)", Env{"x":12, "y":1}, "1729"},
		{"pow(x, 3)+pow(y, 3)", Env{"x":9, "y":10}, "1729"},
		{"5/9*(F-32)", Env{"F":-40}, "-40"},
		{"5/9*(F-32)", Env{"F":32}, "0"},
		{"5/9*(F-32)", Env{"F":212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q want %q\n", test.expr, test.env, test.want)
		}
	}
}