package jsgo

// #include "mujs.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// JsGlobal ...
type JsGlobal struct {
	state *JsState
}

// Set ...
func (global *JsGlobal) Set(name string, value JsValue) {
	property := C.CString(name)
	defer C.free(unsafe.Pointer(property))

	// push value
	C.js_getregistry(global.state.vm, value.reference())

	// set property
	C.js_setglobal(global.state.vm, property)
}

// Get ...
func (global *JsGlobal) Get(name string) JsValue {
	property := C.CString(name)
	defer C.free(unsafe.Pointer(property))

	// get property
	C.js_getglobal(global.state.vm, property)

	// return the property value
	return newJsValue(global.state)
}
