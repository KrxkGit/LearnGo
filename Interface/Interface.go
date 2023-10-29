package Interface

import "fmt"

type Run interface {
	Wait(sec int)
	Go(speed float32)
}

type RunFast interface {
	Run
	HighSpeed(float32)
}

type FootballPlayer struct {
	sec_wait int
	speed    float32
}

// 为了扩展，最好使用同一形式的接收器，否则难以扩充新的接口。

func (fplayer *FootballPlayer) Wait(sec int) {
	fplayer.sec_wait = sec
	fmt.Printf("wait: %d\n", fplayer.sec_wait)
}

func (fplayer *FootballPlayer) Go(speed float32) {
	fplayer.speed = speed
	fmt.Printf("speed: %f\n", fplayer.speed)
}

type FastFootballPlayer struct {
	FootballPlayer
}

func (fplayer *FastFootballPlayer) HighSpeed(speed float32) {
	fplayer.speed = speed
	fmt.Printf("Run Very Fast\n")
}

func Test() {
	var run Run
	footballPlayer := FootballPlayer{
		sec_wait: 3,
		speed:    12,
	}
	fastPlayer := FastFootballPlayer{footballPlayer}
	run = &footballPlayer
	run.Wait(2)
	run.Go(34)

	var fastrun RunFast
	fastrun = &fastPlayer
	fastrun.HighSpeed(23)

	run = &fastPlayer
	run.Wait(2)
}
