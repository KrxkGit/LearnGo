package Routine

import (
	"fmt"
	"runtime"
)

func TestScheduler() {
	runtime.GOMAXPROCS(3) // 设置允许使用的操作系统内核线程数进行调度，不同的值将导致不同的输出效果
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
