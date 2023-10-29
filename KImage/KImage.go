package KImage

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // 利用其init支持 png
	"io"
	"os"
)

func TestImage(out io.Writer) error {
	filePath := "Pic.png"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, kind, err := image.Decode(file)
	if err != nil {
		fmt.Printf("fail!\n")
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
