package Routine

import (
	"fmt"
	"os"
	"time"
)

func TestBroadcast() {
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) { // 复制i，因为 i是循环块共享的，否则将出错
			x := <-done // 等待
			fmt.Printf("Exit normally: %d.\t x:%v\n", i, x)
		}(i)
	}

	os.Stdin.Read(make([]byte, 1))
	close(done) // 广播，使 routine 退出（channel返回零值）
	time.Sleep(1 * time.Second)

	//os.Stdin.Read(make([]byte, 1))
	fmt.Printf("Exit\n")
	//close(done)
}
