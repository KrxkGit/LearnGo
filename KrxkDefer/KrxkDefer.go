package KrxkDefer

import (
	"fmt"
	"time"
)

func TryDefer() {
	// defer 是类似堆栈的结构，先入后出
	defer DelayDo()() //这里是 defer DelayDo 返回的匿名函数,即push func
	defer DelayDo2()  //这里是 defer DelayDo2 函数，即push DelayDo2
	time.Sleep(4 * time.Second)
	fmt.Println("return TryDefer")
}

func DelayDo() func() {
	start := time.Now()
	fmt.Println("Run DelayDo")
	return func() {
		fmt.Printf("%T: %[1]s\n", time.Since(start))
		fmt.Println("return func")
	}
}

func DelayDo2() {
	fmt.Println("DelayDo2")
}

func Test() {
	fmt.Println("Test Start")
	TryDefer()
}
