// hi
package test

// #cgo CFLAGS: -I${SRCDIR}/packaged/include
// #cgo LDFLAGS: -ltesta
// #include <my.h>
import "C"

import (
	_ "github.com/heeus/AlexeyWhite/testCpp/packaged/include"
)
