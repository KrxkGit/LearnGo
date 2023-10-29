package Panic

import "fmt"

func RaisePanic() {
	panic(3)
}

func Test() {
	defer func() {
		if r := recover(); r == 2 {
			fmt.Println("Process the panic")

		} else {
			panic(r)
		}
	}()
	RaisePanic()
}
