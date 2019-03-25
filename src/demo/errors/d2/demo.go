package main

import "fmt"

//使用牛顿法实现一个平方根函数Sqrt, 并使其实现Error接口
//当Sqrt接受到一个负数或者复数时，应当返回一个非nil的错误值

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环，触发一个致命错误，发生栈溢出。
	//可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？
	//因为fmt包在进行输出也会试图匹配错误，e变量通过实现Error接口成了error类型，
	//在fmt.Sprint(e)时就会调用e.Error()来输出错误的字符串信息，也就是上面的代码相当于：fmt.Sprint(e.Error())
	//因此输出时先转化e变量的类型
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) //执行失败时返回类型的零值与一个错误
	}
	const E = 0.000001
	z := float64(1)
	k := float64(0)
	for ; ; z -= (z*z - x) / (2*z) {
		if z - k <= E && z - k >= -E {
			return z, nil //执行成功时返回nil
		}
		//fmt.Println(z)
		k = z
	}
}

func main()  {
	v, err := Sqrt(2)
	fmt.Printf("value: %v, err: %v\n", v, err) //value: 1.4142135623730951, err: <nil>
	v, err = Sqrt(-2)
	fmt.Printf("value: %v, err: %v\n", v, err) //value: 0, err: cannot Sqrt negative number -2
}
