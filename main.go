package main

import (
	"bytes"
	"image"
	"log"
	"syscall/js"

	"github.com/qeesung/image2ascii/convert"
)

func loadImage(this js.Value, args []js.Value) interface{} {
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

	createAsciiImage(sourceImage)

	return nil
}

func createAsciiImage(img image.Image) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 180
	convertOptions.FixedHeight = 60
	convertOptions.Colored = false

	converter := convert.NewImageConverter()
	result := converter.Image2ASCIIString(img, &convertOptions)
	js.Global().Get("document").
		Call("getElementById", "result").
		Set("innerHTML", result)

}

func registerCallbacks() {
	js.Global().Set("loadImage", js.FuncOf(loadImage))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM GO Initialized")

	registerCallbacks()
	<-c
}
