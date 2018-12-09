package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	v := new(distance)
	v.set(1003)
	x := 1023
	fmt.Println(process(v, x))
}

type adder interface {
	add(int) int
}

type distance struct {
	i int
}

func (d *distance) set(i int) {
	d.i = i
}

func (d distance) add(x int) int {
	return d.i + x
}

func process(a adder, x int) int {
	return a.add(x)
}
