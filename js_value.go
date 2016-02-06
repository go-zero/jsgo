package jsgo

/*
#include "mujs.h"

typedef int (*checkfunc)(js_State*, int);

int check(void* f, js_State* state) {
	return ((checkfunc) f)(state, 0);
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// JsValue ...
type JsValue struct {
	state *JsState
	ref   *C.char
}

func newJsValue(state *JsState) *JsValue {
	value := &JsValue{state, C.js_ref(state.vm)}
	runtime.SetFinalizer(value, (*JsValue).unref)
	return value
}

func (value *JsValue) unref() {
	C.js_unref(value.state.vm, value.ref)
}

func (value *JsValue) check(f unsafe.Pointer) bool {
	// put the reference on stack
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)

	// use the 'check' wrapper to call the validation function
	return C.check(f, value.state.vm) == 1
}

// String ...
func (value *JsValue) String() string {
	// put the reference on stack
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)

	toString := C.js_tostring(value.state.vm, 0)
	return C.GoString(toString)
}

// IsDefined ...
func (value *JsValue) IsDefined() bool {
	return value.check(C.js_isdefined)
}

// IsUndefined ...
func (value *JsValue) IsUndefined() bool {
	return value.check(C.js_isundefined)
}

// IsNumber ...
func (value *JsValue) IsNumber() bool {
	return value.check(C.js_isnumber)
}

// IsString ...
func (value *JsValue) IsString() bool {
	return value.check(C.js_isstring)
}

// IsBoolean ...
func (value *JsValue) IsBoolean() bool {
	return value.check(C.js_isboolean)
}

// IsNull ...
func (value *JsValue) IsNull() bool {
	return value.check(C.js_isnull)
}

// IsPrimitive ...
func (value *JsValue) IsPrimitive() bool {
	return value.check(C.js_isprimitive)
}

// IsCallable ...
func (value *JsValue) IsCallable() bool {
	return value.check(C.js_iscallable)
}

// IsObject ...
func (value *JsValue) IsObject() bool {
	return value.check(C.js_isobject)
}

// IsArray ...
func (value *JsValue) IsArray() bool {
	return value.check(C.js_isarray)
}

// IsRegexp ...
func (value *JsValue) IsRegexp() bool {
	return value.check(C.js_isregexp)
}
