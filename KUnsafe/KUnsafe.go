package KUnsafe

import (
	"fmt"
	"unsafe"
)

func TestUnsafe() {
	type kStruct struct {
		a int     // 64bit系统下为 int64，size为8
		b [2]byte // byte size为8
	}

	var x kStruct
	// Sizeof 只统计固定部分
	fmt.Printf("Size: %d\n", unsafe.Sizeof(x))
	fmt.Printf("Align: %d\t%d\n", unsafe.Alignof(x), unsafe.Alignof(x.b))

	// 指针
	x = kStruct{
		a: 1,
		b: [2]byte{'K', 'r'},
	}
	// 应避免保存一个uintptr类型的临时变量，因为 gc 可能会移动内存，但uintptr无法自动被gc更新。应直接转为unsafe.Pointer类型
	ptrA := (*[2]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	ptrA[1] = 'k'
	fmt.Printf("Update value: %d\t%s\n", x.a, x.b)

	// 提示: 不应该这么写! 由于没有指针保存new的结果，gc 可能会直接回收掉
	pkStruct := (*kStruct)(unsafe.Pointer(uintptr(unsafe.Pointer(new(kStruct)))))
	fmt.Printf("Update value: %T\t%[1]v\n", pkStruct)
	pkStruct.a = 1
	fmt.Printf("value: %d", pkStruct.a)
}
