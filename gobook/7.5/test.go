package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	//var w io.Writer
	//fmt.Print(w==nil)

	//var w io.ReadWriter
	////w = os.Stdout
	////rw := w.(io.ReadWriter)
	////fmt.Print(rw)
	////
	////w = rw
	//rw := w.(io.Writer)
	//fmt.Print(rw )

	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success: ok, f == os.Stdout
	b, ok := w.(*bytes.Buffer)

	fmt.Println(f, b, ok, w)
	switch x := w.(type) {
	case nil:

	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
