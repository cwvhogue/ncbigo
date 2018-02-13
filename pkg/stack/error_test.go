// ===========================================================================
//
// File Name:  error_test.go
//
// Author:  Christopher Hogue
//
// ==========================================================================

package stack

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ErrorTesterT struct{}

var _ = Suite(&ErrorTesterT{})

func (s *ErrorTesterT) TestErrors(c *C) {

	terr := ErrorStack(fmt.Errorf("error message %d - a typical error in Go", 1))
	fmt.Printf("Test Error String is: [%s]\n", terr.String())
	fmt.Printf("Test Error Stack is: [%s]\n", terr.Stack())
	fmt.Printf("Test Error All is: [%s]\n", terr.All())
	fmt.Printf("Test Error v is: [%v]\n", terr)
	c.Assert(terr, Not(IsNil))

	terr = ErrorF("Error message %d - another %s with stack trace", 2, "error")
	fmt.Printf("Test Error String is: [%s]\n", terr.String())
	fmt.Printf("Test Error Stack is: [%s]\n", terr.Stack())
	fmt.Printf("Test Error All is: [%s]\n", terr.All())
	fmt.Printf("Test Error v is: [%v]\n", terr)
	c.Assert(terr, Not(IsNil))
}
