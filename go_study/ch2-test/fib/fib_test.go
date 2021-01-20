
package fib

import (
	"testing"
)

var a int

func TestFiblist(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var {
	// 	a int = 1
	// 	b = 1
	// }
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log("", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

//变量赋值
// 与其他语言的差异
// 1.赋值可以进行自动类型推断
// 2.在一个赋值语句中可以对多个变量进行同时赋值

func TestExchange(t *testing.T){
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	a,b = b,a
	t.Log(a,b)

}

