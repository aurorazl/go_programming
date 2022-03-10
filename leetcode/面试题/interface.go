package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var a *int
	var g = new(int)
	fmt.Printf("%p", g)
	var c = new([]int)
	fmt.Println(a, g, c, len(*c), cap(*c))
	fmt.Printf("%p\n%p\n", c, *c)
	var d *[]int
	//d = append(d,1)
	fmt.Println(d, &d)
	fmt.Println(&d, *c, len(*d), cap(*d))
	fmt.Printf("%p", &d)
	fmt.Printf("%p", d)
	var f *[]int
	var h []int
	f = &h
	var j int
	fmt.Println(j)

	//var buf *bytes.Buffer
	var buf = new(bytes.Buffer)
	buf.Write([]byte("hello"))
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
