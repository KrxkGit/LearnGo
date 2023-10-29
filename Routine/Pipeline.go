package Routine

import (
	"fmt"
	"time"
)

func Counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func Square(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x
		time.Sleep(30 * time.Millisecond)
	}
	close(out)
}

func Printer(in <-chan int /*隐式转换为单向channel*/, done chan<- struct{}) {
	for x := range in {
		x = x * x
		fmt.Printf("Value: %d\n", x)
	}
	done <- struct{}{}
}

func TestPipe() {
	defer CalculateTime()()
	c1 := make(chan int)
	c2 := make(chan int)
	done := make(chan struct{})
	go Counter(c1)
	go Square(c1, c2)
	go Printer(c2, done)
	<-done
}
