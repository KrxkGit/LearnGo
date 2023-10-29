package KReflect

import (
	"fmt"
	"reflect"
)

func TestReflect() {
	a := 1
	// 反射读值
	v := reflect.ValueOf(a)
	fmt.Printf("Type: %T\nValue: %[1]v\n", v)

	t := v.Type()
	fmt.Printf("Real Type: %v\nKind: %v\n", t, v.Kind())

	pa := reflect.ValueOf(&a)
	fmt.Printf("Type: %T\nValue: %[1]v\n", pa)

	va := pa.Elem()
	fmt.Printf("Type: %T\nValue: %[1]v\n", va)

	// 反射设置 Value
	fmt.Printf("Can get addr?: %v\n", va.CanAddr())
	if va.CanAddr() {
		px := va.Addr().Interface().(*int)
		*px = 2
		fmt.Printf("New value: {x:%d, a:%d}\n", *px, a)
	}

	// 反射设置值的另一种方式
	b := 1
	pb := reflect.ValueOf(&b).Elem()
	if pb.CanAddr() {
		pb.Set(reflect.ValueOf(3))
		fmt.Printf("New value: {b:%d}\n", b)
	}

	// 反射还有许多功能
}
