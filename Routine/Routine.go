package Routine

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func CalculateTime() func() int {
	start := time.Now()
	return func() int {
		fmt.Printf("\nCostTime: %v", time.Since(start))
		return 1
	}
}

func Test() {
	defer CalculateTime()()
	go spinner(100 * time.Millisecond)
	time.Sleep(5 * time.Second)

}
