package Interface

import (
	"fmt"
)

func TestAssert() {
	var run Run
	fplayer := FastFootballPlayer{FootballPlayer{
		sec_wait: 12,
		speed:    14,
	}}
	//fplayer1 := FootballPlayer{
	//	1,
	//	14,
	//}
	run = &fplayer
	fastrun := run.(RunFast) // 无法提升会 panic
	fmt.Printf("%T", fastrun)
	if fastrun, ok := run.(RunFast); ok { // 该方法不会引发 panic
		fastrun.HighSpeed(20)
		fmt.Printf("Promote successfully\n")
	} else {
		fmt.Printf("No Promotion\n")
	}

	switch fastrun.(type) { // 类型 switch
	case RunFast:
		fmt.Println("FastRun match")
	case Run:
		fmt.Println("Run match")
	default:
		panic("No rule match")
	}
}
