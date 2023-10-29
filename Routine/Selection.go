package Routine

import (
	"fmt"
	"os"
	"time"
)

func TestSelection() {
	abort := make(chan struct{})
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

Loop: // label 该循环
	for i := 1; i <= 10; i++ {
		select {
		case <-ticker.C:
			fmt.Printf("Tick: %d\n", i)
		case <-abort:
			fmt.Printf("Recieve Abort Command.\n")
			break Loop // 跳出外层for循环
			//default:
			//	fmt.Printf("default select\n")
			//	//do nothing
		}
	}

	ticker.Stop()
	fmt.Printf("Launch!\n")
}
