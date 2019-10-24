package dynamic

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>

void call_func(void *p, void *args) {
	void (*func)(void *) = p;
	func(args);
}
*/
import (
	"C"
)
import (
	"fmt"
	"unsafe"

	"github.com/ophum/kudryavka/waf/gate"
)

type DynamicGate struct {
	path string
}

func NewDynamicGate(path string) (*DynamicGate, error) {
	dg := &DynamicGate{
		path: path,
	}

	return dg, nil
}

func (d *DynamicGate) Check(args gate.CheckList) error {
	handle := C.dlopen(C.CString(d.path), C.RTLD_LAZY)
	if handle == nil {
		return fmt.Errorf("Error: failed to open shared object.\n")
	}
	defer C.dlclose(handle)

	check_func := C.dlsym(handle, C.CString("check"))
	if check_func == nil {
		return fmt.Errorf("Error: could not found check function pointer.\n")
	}
	C.call_func(check_func, unsafe.Pointer(&args))
	return nil
}
