package jsgo

// #include "mujs.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// JsState ...
type JsState struct {
	vm *C.js_State
}

// NewJsState ...
func NewJsState() *JsState {
	state := new(JsState)
	state.vm = C.js_newstate(nil, nil, 1) // JS_STRICT
	return state
}

// Free ...
func (state *JsState) Free() {
	C.js_freestate(state.vm)
}

// DoString ...
func (state *JsState) DoString(text string) (JsValue, error) {
	source := C.CString(text)
	defer C.free(unsafe.Pointer(source))

	if rc := C.js_ploadstring(state.vm, C.CString("[string]"), source); rc != 0 {
		return nil, state.getError()
	}

	C.js_pushglobal(state.vm)

	if rc := C.js_pcall(state.vm, 0); rc != 0 {
		return nil, state.getError()
	}

	return newJsValue(state), nil
}

func (state *JsState) getError() error {
	// put all 3 error attributes on stack
	C.js_getproperty(state.vm, 0, C.CString("name"))
	C.js_getproperty(state.vm, 0, C.CString("message"))
	C.js_getproperty(state.vm, 0, C.CString("stackTrace"))

	// grab the values in the same order we put on stack
	name := C.js_tostring(state.vm, 1)
	message := C.js_tostring(state.vm, 2)
	stackTrace := C.js_tostring(state.vm, 3)

	// pop the error, name, message and stackTrace
	C.js_pop(state.vm, 4)

	// return the new error
	return &JsError{C.GoString(name), C.GoString(message), C.GoString(stackTrace)}
}
