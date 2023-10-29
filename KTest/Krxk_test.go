package KTest

import (
	"HelloWorld/KImage"
	"fmt"
	"os"
	"testing"
)

func TestKrxk(t *testing.T) {
	file, err := os.Create("d:/1.jpg")
	if err != nil {
		t.Error("Krxk Test Error")
	}
	fmt.Printf("Exit due to fail\n")
	defer file.Close()
	err = KImage.TestImage(file)
	if err != nil {
		t.Error("Convert Error")

	}
	fmt.Printf("Exit due to fail2\n")
}
