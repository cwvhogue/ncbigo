// ===========================================================================
//
// File Name:  error.go
//
// Authors:  Andrew Hume, Christopher Hogue
//
// ==========================================================================

// Expands the Go error type to include a stack trace.

// For the typical Go error return value,
// up-convert it to include a stack trace as follows:
// err := ErrorStack(fmt.Errorf("error message %d - a typical error in Go", 1))
//
// Or you can create your own error with stack trace and fmt.Sprintf
// style formatting.
//
// err = ErrorF("error message %d - another %s with stack trace", 2, "error")

package stack

import (
	"fmt"

	"github.com/go-stack/stack"
)

// ErrorT - modified error type with stack trace
type ErrorT struct {
	Err string
	Stk string
}

// ErrorStack - returns the modified error
func ErrorStack(err error) *ErrorT {
	if err == nil {
		return nil
	}
	error := ErrorT{Err: err.Error(), Stk: CallStack()}
	return &error
}

// ErrorS - returns an ErrorT from a string
func ErrorS(s string) *ErrorT {
	if s == "" {
		return nil
	}
	error := ErrorT{Err: s, Stk: CallStack()}
	return &error
}

// ErrorF - returns an ErrorT with a fmt.Sprintf wrapper for format arguments
func ErrorF(f string, a ...interface{}) *ErrorT {
	error := ErrorT{Err: fmt.Sprintf(f, a...), Stk: CallStack()}
	return &error
}

// err.String() returns the string portion
func (e *ErrorT) String() string {
	if e == nil {
		return ""
	}
	return e.Err
}

// Stack - method returns the stack portion as err.Stack()
func (e *ErrorT) Stack() string {
	if e == nil {
		return ""
	}
	return e.Stk
}

// All - method returns error string and stack as err.All()
func (e *ErrorT) All() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s\n%s", e.Err, e.Stk)
}

// Error - method implements the Go error interface as err.Error()
func (e *ErrorT) Error() string {
	return e.String()
}

// Assert - implements a classic assert function
func Assert(val bool) {
	if !val {
		panic(nil)
	}
}

// CallStack - returns a nice compressed call stack.
func CallStack() string {
	myfn := fmt.Sprintf // get around go vet enthusiasm
	ret := ""
	stack := stack.Trace()
	stack = stack[1:]
	// trim the annoying goexit at top
	n := len(stack)
	if myfn("%n", stack[n-1]) == "goexit" {
		stack = stack[:n-1]
	}
	for _, e := range stack {
		ret += myfn("<-%+n(%s:%d)", e, e, e)
	}
	return ret[2:]
}
