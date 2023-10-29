package Constant

import (
	"fmt"
	"time"
)

const (
	Ipv4                   = 4
	IP_local string        = "192.168.0.101"
	nolay    time.Duration = 0
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func init() {
	fmt.Println(Ipv4)
	fmt.Println(IP_local)
	fmt.Printf("%T:%[1]v\n", nolay)
}

const (
	YiB = 10 << 100
	ZiB = 10 << 90
)

func Krxk() {
	fmt.Println(Sunday)
	fmt.Println(Wednesday)
	fmt.Println(Friday)
	fmt.Println(FlagUp)
	fmt.Println(YiB / ZiB) // "1024"

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}
