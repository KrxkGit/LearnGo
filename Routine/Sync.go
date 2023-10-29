package Routine

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	//mu  sync.Mutex
	mu  sync.RWMutex
	obj krxkObj
	wg  sync.WaitGroup
)

type krxkObj struct {
	// 只要 val 与 no 保持一致即达到目标
	val int
	no  int
}

func (k *krxkObj) increaseObj(ins int) {
	k.val += ins
	time.Sleep(3 * time.Millisecond) // 加入延时，使效果更明显
	k.no += ins
}

func WriteObj() {
	mu.Lock()
	defer mu.Unlock()
	obj.increaseObj(1)
	wg.Done()
}

func ReadObj() {
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(time.Duration(math.Mod(float64(rand.Int()), 10)) * time.Millisecond) // 加入延时，使效果更明显
	fmt.Printf("Value: %v\n", obj)

	wg.Done()
}

func TestSync() {
	defer CalculateTime()()
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go WriteObj()
		go ReadObj()
	}

	wg.Wait()
}
