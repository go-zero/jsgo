package jsgo

import "fmt"

// JsError ...
type JsError struct {
	Name       string
	Message    string
	StackTrace string
}

// JsError ...
func (err JsError) Error() string {
	return fmt.Sprintf("%s: %s%s", err.Name, err.Message, err.StackTrace)
}
