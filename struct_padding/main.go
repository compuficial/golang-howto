package main

import (
	"fmt"
	"unsafe"
)

type Inefficient struct {
	field1 bool
	field2 int64
	field3 bool
	field4 int64
}

type Efficient struct {
	field2 int64
	field4 int64
	field1 bool
	field3 bool
}

func main() {
	ineff := unsafe.Sizeof(Inefficient{})
	eff := unsafe.Sizeof(Efficient{})

	fmt.Printf("Inefficient struct size: %d\n", ineff)
	fmt.Printf("Efficient struct size: %d\n", eff)
}
