// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dig

import (
	"fmt"
	"io"

	"go.uber.org/dig/internal/digreflect"
	"go.uber.org/dig/internal/dot"
)

// Error is an interface implemented by all Dig errors.
//
// Use this interface, in conjunction with [RootCause], in order to
// determine if errors you encounter come from Dig, or if they come
// from provided constructors or invoked functions. See [RootCause]
// for more info.
type Error interface {
	error

	// Writes the message or context for this error in the chain.
	//
	// Note: the Error interface must always have a private function
	// such as this one in order to maintain properly sealed.
	//
	// verb is either %v or %+v.
	writeMessage(w io.Writer, v string)
}

// a digError is a dig.Error with additional functionality for
// internal use - namely the ability to be formatted.
type digError interface {
	Error
	fmt.Formatter
}

// A PanicError occurs when a panic occurs while running functions given to the container
// with the [RecoverFromPanic] option being set. It contains the panic message from the
// original panic. A PanicError does not wrap other errors, and it does not implement
// dig.Error, meaning it will be returned from [RootCause]. With the [RecoverFromPanic]
// option set, a panic can be distinguished from dig errors and errors from provided/
// invoked/decorated functions like so:
//
//	rootCause := dig.RootCause(err)
//
//	var pe dig.PanicError
//	var de dig.Error
//	if errors.As(rootCause, &pe) {
//		// This is caused by a panic
//	} else if errors.As(err, &de) {
//		// This is a dig error
//	} else {
//		// This is an error from one of my provided/invoked functions or decorators
//	}
//
// Or, if only interested in distinguishing panics from errors:
//
//	var pe dig.PanicError
//	if errors.As(err, &pe) {
//		// This is caused by a panic
//	} else {
//		// This is an error
//	}
type PanicError struct {
	// The function the panic occurred at
	fn *digreflect.Func

	// The panic that was returned from recover()
	Panic any
}

// Format will format the PanicError, expanding the corresponding function if in +v mode.
func (e PanicError) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (e PanicError) Error() string { _ = "STUB: not implemented"; return "" }

// formatError will call a dig.Error's writeMessage() method to print the error message
// and then will automatically attempt to print errors wrapped underneath (which can create
// a recursive effect if the wrapped error's Format() method then points back to this function).
func formatError(e digError, w fmt.State, v rune) { _ = "STUB: not implemented"; return }

// "context: " or "context:\n"

// Will route back to this function recursively if next error
// is also wrapped and points back here

// RootCause returns the first non-dig.Error in a chain of wrapped
// errors, if there is one. Otherwise, RootCause returns the error
// on the bottom of the chain of wrapped errors.
//
// Use this function and errors.As to differentiate between Dig errors
// and errors thrown by provided constructors or invoked functions:
//
//	rootCause := dig.RootCause(err)
//	var de dig.Error
//	if errors.As(rootCause, &de) {
//	    // Is a Dig error
//	} else {
//	    // Is an error thrown by one of my provided/invoked/decorated functions
//	}
//
// See [PanicError] for an example showing how to additionally detect
// and handle panics in provided/invoked/decorated functions.
func RootCause(err error) error {
	_ = "STUB: not implemented"

	// Dig down to first non dig.Error, or bottom of chain
	return nil
}

// errInvalidInput is returned whenever the user provides bad input when
// interacting with the container. May optionally have a more detailed
// error wrapped underneath.
type errInvalidInput struct {
	Message string
	Cause   error
}

var _ digError = errInvalidInput{}

// newErrInvalidInput creates a new errInvalidInput, wrapping the given
// other error that caused this error. If there is no underlying cause,
// pass in nil. This will cause all attempts to unwrap this error to return
// nil, replicating errors.Unwrap's behavior when passed an error without
// an Unwrap() method.
func newErrInvalidInput(msg string, cause error) errInvalidInput {
	_ = "STUB: not implemented"
	return *new(errInvalidInput)
}

func (e errInvalidInput) Error() string { _ = "STUB: not implemented"; return "" }

func (e errInvalidInput) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errInvalidInput) writeMessage(w io.Writer, _ string) { _ = "STUB: not implemented"; return }

func (e errInvalidInput) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

// errProvide is returned when a constructor could not be Provided into the
// container.
type errProvide struct {
	Func   *digreflect.Func
	Reason error
}

var _ digError = errProvide{}

func (e errProvide) Error() string { _ = "STUB: not implemented"; return "" }

func (e errProvide) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errProvide) writeMessage(w io.Writer, verb string) { _ = "STUB: not implemented"; return }

func (e errProvide) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

// errConstructorFailed is returned when a user-provided constructor failed
// with a non-nil error.
type errConstructorFailed struct {
	Func   *digreflect.Func
	Reason error
}

var _ digError = errConstructorFailed{}

func (e errConstructorFailed) Error() string { _ = "STUB: not implemented"; return "" }

func (e errConstructorFailed) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errConstructorFailed) writeMessage(w io.Writer, verb string) {
	_ = "STUB: not implemented"
	return
}

func (e errConstructorFailed) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

// errArgumentsFailed is returned when a function could not be run because one
// of its dependencies failed to build for any reason.
type errArgumentsFailed struct {
	Func   *digreflect.Func
	Reason error
}

var _ digError = errArgumentsFailed{}

func (e errArgumentsFailed) Error() string { _ = "STUB: not implemented"; return "" }

func (e errArgumentsFailed) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errArgumentsFailed) writeMessage(w io.Writer, verb string) {
	_ = "STUB: not implemented"
	return
}

func (e errArgumentsFailed) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

// errMissingDependencies is returned when the dependencies of a function are
// not available in the container.
type errMissingDependencies struct {
	Func   *digreflect.Func
	Reason error
}

var _ digError = errMissingDependencies{}

func (e errMissingDependencies) Error() string { _ = "STUB: not implemented"; return "" }

func (e errMissingDependencies) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errMissingDependencies) writeMessage(w io.Writer, verb string) {
	_ = "STUB: not implemented"
	return
}

func (e errMissingDependencies) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

// errParamSingleFailed is returned when a paramSingle could not be built.
type errParamSingleFailed struct {
	Key    key
	Reason error
	CtorID dot.CtorID
}

var _ digError = errParamSingleFailed{}

func (e errParamSingleFailed) Error() string { _ = "STUB: not implemented"; return "" }

func (e errParamSingleFailed) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errParamSingleFailed) writeMessage(w io.Writer, _ string) {
	_ = "STUB: not implemented"
	return
}

func (e errParamSingleFailed) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (e errParamSingleFailed) updateGraph(g *dot.Graph) { _ = "STUB: not implemented"; return }

// errParamGroupFailed is returned when a value group cannot be built because
// any of the values in the group failed to build.
type errParamGroupFailed struct {
	Key    key
	Reason error
	CtorID dot.CtorID
}

var _ digError = errParamGroupFailed{}

func (e errParamGroupFailed) Error() string { _ = "STUB: not implemented"; return "" }

func (e errParamGroupFailed) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errParamGroupFailed) writeMessage(w io.Writer, _ string) { _ = "STUB: not implemented"; return }

func (e errParamGroupFailed) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (e errParamGroupFailed) updateGraph(g *dot.Graph) { _ = "STUB: not implemented"; return }

// missingType holds information about a type that was missing in the
// container.
type missingType struct {
	Key key // item that was missing

	// If non-empty, we will include suggestions for what the user may have
	// meant.
	suggestions []key
}

// Format prints a string representation of missingType.
//
// With %v, it prints a short representation ideal for an itemized list.
//
//	io.Writer
//	io.Writer: did you mean *bytes.Buffer?
//	io.Writer: did you mean *bytes.Buffer, or *os.File?
//
// With %+v, it prints a longer representation ideal for standalone output.
//
//	io.Writer: did you mean to Provide it?
//	io.Writer: did you mean to use *bytes.Buffer?
//	io.Writer: did you mean to use one of *bytes.Buffer, or *os.File?
func (mt missingType) Format(w fmt.State, v rune) { _ = "STUB: not implemented"; return }

// errMissingType is returned when one or more values that were expected in
// the container were not available.
//
// Multiple instances of this error may be merged together by appending them.
type errMissingTypes []missingType // inv: len > 0

var _ digError = errMissingTypes(nil)

func newErrMissingTypes(c containerStore, k key) errMissingTypes {
	_ = "STUB: not implemented"
	// Possible types we will look for in the container. We will always look
	// for pointers to the requested type and some extras on a per-Kind basis.
	return *new(errMissingTypes)
}

// The user requested a pointer but maybe we have a value.

// Maybe the user meant a slice of pointers while we have the slice of elements

// Maybe the user meant a slice of elements while we have the slice of pointers

// Maybe the user meant an array of pointers while we have the array of elements

// Maybe the user meant an array of elements while we have the array of pointers

// Maybe we have an implementation of the interface.

// Maybe we have an interface that this type implements.

// range through c.providers is non-deterministic. Let's sort the list of
// suggestions.

func (e errMissingTypes) Error() string { _ = "STUB: not implemented"; return "" }

func (e errMissingTypes) writeMessage(w io.Writer, v string) { _ = "STUB: not implemented"; return }

// With %v, we need a space between : since the error
// won't be on a new line.

func (e errMissingTypes) Format(w fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (e errMissingTypes) updateGraph(g *dot.Graph) { _ = "STUB: not implemented"; return }

type errVisualizer interface {
	updateGraph(*dot.Graph)
}
