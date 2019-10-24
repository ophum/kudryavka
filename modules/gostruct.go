package main

import (
	"C"
	"fmt"
	"unsafe"

	"github.com/ophum/kudryavka/waf/gate"
)

func main() {}

//export check
func check(tp unsafe.Pointer) {
	t := (*gate.CheckList)(tp)
	fmt.Println(t)
}
