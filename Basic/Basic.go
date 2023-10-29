package Basic

import (
	"HelloWorld/Constant"
	"bytes"
	"fmt"
	"math"
	"unicode/utf8"
)

// var KrxkUint uint
const (
	krxk16     = "\xe4\xb8\x96\xe7\x95\x8c"
	krxk8      = "\024"
	krxkOrigin = `Hello\\World`
)

func init() {
	//KrxkUint = 11529215046068469760
	fmt.Println("Init")
}

func Print() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

}

func PrintFormat() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

	fmt.Printf("%[2]f\t%#[1]f\n", math.MaxFloat64, math.MaxFloat32)

	fmt.Println(1i * 1i) // "(-1+0i)", i^2 = -1

	fmt.Printf("%s %s %s  %d\n", krxk8, krxk16, krxkOrigin, len(krxk8))

}

func PrintUnicode() {
	s := "hello, world"
	fmt.Println(len(s))               // "12"
	fmt.Printf("%c %c\n", s[0], s[5]) // "104 119" ('h' and 'w')
	//var c = s[100]
	fmt.Println(s[1:5])
}

func PrintDifferentWordLen() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range s {
		n++
	}
	fmt.Printf("Count:%d\n", n)
}

func PrintRune() {
	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r)) // "プログラム"

}

// IntsToString is like fmt.Sprint(values) but adds commas.
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	buf.WriteString("世界")

	//t := "你好"
	var d int
	fmt.Scanf("%d", &d)
	//r := []rune(t)
	//buf.WriteRune(r[0])
	//buf.WriteString(strconv.Itoa(d))
	//buf.WriteString(strconv.Itoa(d))
	q := fmt.Sprintf(" %d", d)
	buf.WriteString(q)
	Constant.Krxk()
	//fmt.Println(strconv.ParseInt("1246", 10, 64))
	return buf.String()
}
