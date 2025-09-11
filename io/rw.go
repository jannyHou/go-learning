package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func IOTest() {
	fmt.Println("*****  I/O tests  *****")
	var buf bytes.Buffer
	buf.WriteString("Hello")
	fmt.Printf("buf is %v, buf to string is %v\n", buf, buf.String())

	fmt.Fprintf(&buf, " world")
	fmt.Printf("buf is %v, buf to string is %v\n", buf, buf.String())

	r := bytes.NewReader([]byte("Hello World!"))
	data, _ := io.ReadAll(r)
	fmt.Printf("data is %v, with type %T\n", data, data)

	// io.Copy copies source reader to destination writer
	r1 := bytes.NewReader([]byte("Foo\n"))
	io.Copy(os.Stdout, r1)

	// fmt.Fprintln prints string to destination writer
	r2 := strings.NewReader("Bar")
	data2, _ := io.ReadAll(r2)
	fmt.Fprintln(os.Stdout, string(data2))
}
