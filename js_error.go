package jsgo

// #include "mujs.h"
import "C"

// JsError ...
type JsError interface {
	error
	Value() JsValue
}

type basicError struct {
	value JsValue
}

func newJsError(state *JsState) JsError {
	return &basicError{newJsValue(state)}
}

// Value ...
func (err basicError) Value() JsValue {
	return err.value
}

// Error ...
func (err basicError) Error() string {
	return err.value.String()
}
