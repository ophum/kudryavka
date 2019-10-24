package main

import "C"
import (
	"fmt"
	"unsafe"

	"github.com/ophum/kudryavka/waf/gate"
)

//export check
func check(args unsafe.Pointer) {
	req := (*gate.CheckList)(args)
	fmt.Println(req)
}

func main() {

}
