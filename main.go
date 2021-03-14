package main

import (
	"bytes"
	"image"
	"log"
	"syscall/js"

	"github.com/qeesung/image2ascii/convert"
)

func createAsciiImage(this js.Value, args []js.Value) interface{} {
	arr := args[0]
	buf := make([]uint8, arr.Get("byteLength").Int())

	js.CopyBytesToGo(buf, arr)

	reader := bytes.NewReader(buf)

	var err error
	sourceImage, _, err := image.Decode(reader)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return getConvertedAscii(sourceImage)
}

func getConvertedAscii(img image.Image) string {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 180
	convertOptions.FixedHeight = 60
	convertOptions.Colored = false

	converter := convert.NewImageConverter()
	return converter.Image2ASCIIString(img, &convertOptions)
}

func registerCallbacks() {
	js.Global().Set("createAsciiImage", js.FuncOf(createAsciiImage))
}

func main() {
	c := make(chan bool, 0)

	println("WASM GO Initialized")

	registerCallbacks()
	<-c
}
