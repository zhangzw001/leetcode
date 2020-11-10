package main

import "fmt"

//定义一个接口,存在Value方法
type emptyInterface interface {
	Value(key interface{}) interface{}
}

// 定义一个empty类型
type empty int

// 定义empty类型的Value方法,实现emptyInterface
func (e *empty) Value(key interface{}) interface{} {
	return key
}


func main() {
	// 定义一个empty类型结构体
	e := new(empty)
	// 定义一个emptyInterface类型接口
	var eIface emptyInterface
	// 接口赋值,完成方法的重写
	eIface = e
	// 对接口进行调用
	fmt.Println(eIface.Value(1))
	// 显然通过结构体也能调用, 但是一般我们不这么做
	fmt.Println(e.Value(1))
}

