package jsgo

// #include "mujs.h"
// #include <stdlib.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// JsState ...
type JsState struct {
	vm     *C.js_State
	global *JsGlobal
}

// NewJsState ...
func NewJsState() *JsState {
	state := new(JsState)
	state.vm = C.js_newstate(nil, nil, 1) // JS_STRICT
	state.global = &JsGlobal{state}
	runtime.SetFinalizer(state, (*JsState).free)
	return state
}

func (state *JsState) free() {
	C.js_freestate(state.vm)
}

// DoString ...
func (state *JsState) DoString(text string) (JsValue, JsError) {
	source := C.CString(text)
	defer C.free(unsafe.Pointer(source))

	if rc := C.js_ploadstring(state.vm, C.CString("[string]"), source); rc != 0 {
		return nil, newJsError(state)
	}

	C.js_pushglobal(state.vm)

	if rc := C.js_pcall(state.vm, 0); rc != 0 {
		return nil, newJsError(state)
	}

	return newJsValue(state), nil
}

// Global ...
func (state *JsState) Global() *JsGlobal {
	return state.global
}
