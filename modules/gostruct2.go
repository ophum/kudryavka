package main

import (
	"C"
	"unsafe"
	"fmt"
	"github.com/ophum/kudryavka/test"
)

func main() {}


//export print_test
func print_test(tp unsafe.Pointer) {
	t := (*test.Test)(tp)
	fmt.Println("名前は", t.Name, "年は", t.Age)
}
