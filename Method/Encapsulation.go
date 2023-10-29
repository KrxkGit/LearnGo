package Method

import (
	"fmt"
)

type Encap struct {
	Krxkint int
	KrxkBike
}

type KrxkBike struct {
	wheel_count int
}

func (k KrxkBike) CountWheel() int {
	return k.wheel_count
}

func TestEncap() {
	encap := Encap{
		Krxkint:  0,
		KrxkBike: KrxkBike{2},
	}
	fmt.Printf("数量： %d\n", encap.CountWheel())

	f := KrxkBike.CountWheel
	bike := KrxkBike{4}
	fmt.Printf("%T\n%d\n", f, f(bike))
}
