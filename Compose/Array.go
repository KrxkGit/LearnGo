package Compose

import (
	"crypto/sha256"
	"fmt"
)

var a = [3]int{1, 2, 3}
var q = [...]int{1, 2, 5, 6}

func KrxkArray() {
	CompareSHA256()
	SizeTest()
	fmt.Println(CompareArray())
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Printf("Len:%d\tVaule:%d", len(q), q)
	PrintArray()
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func PrintArray() {
	// 根据索引定义
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	fmt.Println(symbol)

	r := [...]int{99: -1}
	r1 := [...]int{88: -1, 77: 123}
	fmt.Println(r)
	fmt.Println(r1)
}

func CompareArray() bool {
	// 必须长度相等，长度不等视为不同类型
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	return arr1 == arr2
}

func CompareSHA256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	c3 := c2
	Ones(&c3, c2)
	fmt.Printf("c1: %x\nc2:%x\nEqual:%t\nType:%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(c3)
}

func Ones(ptr *[32]uint8, tempArr [32]uint8) {
	for i := range ptr {
		ptr[i] = 1
	}
	tempArr[1] = 2
	fmt.Printf("TempArr: %x\n", tempArr)
	s := make([]uint, 10)[:2]
	s[1] = 2
	fmt.Printf("%d\n", s)
}

func SizeTest() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i) //append 函数返回的结果不一定引用相同的底层数组，故需要赋值（与Python不同）
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
