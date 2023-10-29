package Args

import "fmt"

func TestArgs() {
	display("%.3f\t%d", 0.2, 3)
}

func display(format string, args ...interface{}) {
	fmt.Printf("Krxk: "+format, args...)
}
