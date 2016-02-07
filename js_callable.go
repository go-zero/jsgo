package jsgo

// #include "mujs.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// JsCallable ...
type JsCallable struct {
	*JsObject
}

// Call ...
func (callable *JsCallable) Call(args ...interface{}) (JsValue, error) {
	// push function
	C.js_getregistry(callable.state.vm, callable.ref)

	// push the this value to be used by the function
	C.js_pushundefined(callable.state.vm) // for now, just pushing undefined

	// push the arguments
	argsCount := 0
	for _, arg := range args {
		if arg == nil {
			C.js_pushnull(callable.state.vm)
		} else {
			switch arg := arg.(type) {
			case int:
				C.js_pushnumber(callable.state.vm, C.double(arg))
			case float64:
				C.js_pushnumber(callable.state.vm, C.double(arg))
			case bool:
				value := 0
				if arg {
					value = 1
				}
				C.js_pushboolean(callable.state.vm, C.int(value))
			case string:
				value := C.CString(arg)
				defer C.free(unsafe.Pointer(value))
				C.js_pushstring(callable.state.vm, value)
			}
		}
		argsCount++
	}

	// call the function
	if rc := C.js_pcall(callable.state.vm, C.int(argsCount)); rc != 0 {
		return nil, callable.state.getError()
	}

	// return the result
	return newJsValue(callable.state), nil
}
