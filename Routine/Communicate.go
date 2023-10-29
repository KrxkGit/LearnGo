package Routine

import (
	"fmt"
	"time"
)

var prepare chan int // 全局 channel
var done chan struct{}

func Do(isSend bool) {
	if isSend {
		// 发送
		x := 5
		prepare <- x
		//close(prepare)
		fmt.Printf("已发送：%d\n", x)
		time.Sleep(3 * time.Second)
		//prepare <- x

	} else {
		// 接收
		//var x int
		x := <-prepare
		fmt.Printf("Recv Value: %d\n", x)

		close(prepare)     // 关闭channel，不可再发送
		done <- struct{}{} // 用于让主线程结束等待
	}
}

func TestMulti() chan struct{} {
	prepare = make(chan int)
	done = make(chan struct{})
	go Do(true)
	go Do(false)
	return done
	//<-prepare // 等待
}
