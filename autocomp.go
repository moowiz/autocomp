package autocomp

import (
	"runtime"
)

func KeepFresh() {
	a, b, c, d = runtime.Caller(1)
	fmt.Printf("%v %v %v %v", a, b, c, d)
}
