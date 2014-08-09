package cgo_sample

import "C"
import "fmt"

func Foo() {
	s := C.GoString(C.CString("foo"))
	fmt.Println(s)
}
