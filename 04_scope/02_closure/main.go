package main

import "fmt"

func main() {
	f := test(123)
	f()

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
func test(x int) func() {
	println(&x)
	return func() {
		println(&x, x)
	}
}
func wrapper() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//test函数或wrapper函数返回的是 所定义的匿名函数 和  所引用的环境变量的指针
