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
type JsValue interface {
	Float() float64
	Integer() int
	String() string
	Bool() bool
	IsDefined() bool
	IsUndefined() bool
	IsNumber() bool
	IsString() bool
	IsBool() bool
	IsNull() bool
	IsPrimitive() bool
	IsCallable() bool
	IsObject() bool
	IsArray() bool
	IsRegExp() bool
	unref()
}

type basicValue struct {
	state *JsState
	ref   *C.char
}

func newJsValue(state *JsState) JsValue {
	basic := &basicValue{state, C.js_ref(state.vm)}
	var value JsValue = basic

	if basic.IsObject() {
		object := &JsObject{basic}

		switch {
		case object.IsCallable():
			value = &JsCallable{object}
		case object.IsArray():
			value = &JsArray{object}
		case object.IsRegExp():
			value = &JsRegExp{object}
		default:
			value = object
		}
	}

	runtime.SetFinalizer(value, JsValue.unref)
	return value
}

func (value *basicValue) unref() {
	C.js_unref(value.state.vm, value.ref)
}

// Float ...
func (value *basicValue) Float() float64 {
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)
	return float64(C.js_tonumber(value.state.vm, 0))
}

// Integer ...
func (value *basicValue) Integer() int {
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)
	return int(C.js_toint32(value.state.vm, 0))
}

// String ...
func (value *basicValue) String() string {
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)
	return C.GoString(C.js_tostring(value.state.vm, 0))
}

// Bool ...
func (value *basicValue) Bool() bool {
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)
	return C.js_toboolean(value.state.vm, 0) == 1
}

func (value *basicValue) check(f unsafe.Pointer) bool {
	// put the reference on stack
	C.js_getregistry(value.state.vm, value.ref)
	defer C.js_pop(value.state.vm, 1)

	// use the 'check' wrapper to call the validation function
	return C.check(f, value.state.vm) == 1
}

// IsDefined ...
func (value *basicValue) IsDefined() bool {
	return value.check(C.js_isdefined)
}

// IsUndefined ...
func (value *basicValue) IsUndefined() bool {
	return value.check(C.js_isundefined)
}

// IsNumber ...
func (value *basicValue) IsNumber() bool {
	return value.check(C.js_isnumber)
}

// IsString ...
func (value *basicValue) IsString() bool {
	return value.check(C.js_isstring)
}

// IsBool ...
func (value *basicValue) IsBool() bool {
	return value.check(C.js_isboolean)
}

// IsNull ...
func (value *basicValue) IsNull() bool {
	return value.check(C.js_isnull)
}

// IsPrimitive ...
func (value *basicValue) IsPrimitive() bool {
	return value.check(C.js_isprimitive)
}

// IsCallable ...
func (value *basicValue) IsCallable() bool {
	return value.check(C.js_iscallable)
}

// IsObject ...
func (value *basicValue) IsObject() bool {
	return value.check(C.js_isobject)
}

// IsArray ...
func (value *basicValue) IsArray() bool {
	return value.check(C.js_isarray)
}

// IsRegexp ...
func (value *basicValue) IsRegExp() bool {
	return value.check(C.js_isregexp)
}
