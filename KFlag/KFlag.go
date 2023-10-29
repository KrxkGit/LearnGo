package KFlag

import (
	"flag"
	"fmt"
)

func SetFlag() {
	n := flag.Bool("Run", false, "Set Var Run true or false")
	s := flag.String("NewString", "", "Input a String")

	flag.Parse()
	fmt.Printf("Flag:\n %T\t%[1]v\n%T\t%[2]v", *n, *s)
}
