//
package test

// #include <my.h>
import "C"

func sum() {
	return
}

func helloC() (res int) {
	return C.C_CONST
}
