// Copyright (c) 2021 Uber Technologies, Inc.
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

// An InvokeOption modifies the default behavior of Invoke.
type InvokeOption interface {
	applyInvokeOption(*invokeOptions)
}

type invokeOptions struct {
	Info *InvokeInfo
}

// InvokeInfo provides information about an Invoke.
type InvokeInfo struct {
	Inputs []*Input
}

// FillInvokeInfo is an InvokeOption that writes information on the types
// accepted by the Invoke function into the specified InvokeInfo.
// For example:
//
//			var info dig.InvokeInfo
//			err := c.Invoke(func(string, int){}, dig.FillInvokeInfo(&info))
//
//	  info.Inputs[0].String() will be string.
//	  info.Inputs[1].String() will be int.
func FillInvokeInfo(info *InvokeInfo) InvokeOption {
	_ = "STUB: not implemented"
	return *new(InvokeOption)
}

type fillInvokeInfoOption struct {
	info *InvokeInfo
}

func (o fillInvokeInfoOption) String() string { _ = "STUB: not implemented"; return "" }

func (o fillInvokeInfoOption) applyInvokeOption(opts *invokeOptions) {
	_ = "STUB: not implemented"
	return

	// Invoke runs the given function after instantiating its dependencies.
	//
	// Any arguments that the function has are treated as its dependencies. The
	// dependencies are instantiated in an unspecified order along with any
	// dependencies that they might have.
	//
	// The function may return an error to indicate failure. The error will be
	// returned to the caller as-is.
	//
	// If the [RecoverFromPanics] option was given to the container and a panic
	// occurs when invoking, a [PanicError] with the panic contained will be
	// returned. See [PanicError] for more info.
}

func (c *Container) Invoke(function interface{}, opts ...InvokeOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Invoke runs the given function after instantiating its dependencies.
//
// Any arguments that the function has are treated as its dependencies. The
// dependencies are instantiated in an unspecified order along with any
// dependencies that they might have.
//
// The function may return an error to indicate failure. The error will be
// returned to the caller as-is.
func (s *Scope) Invoke(function interface{}, opts ...InvokeOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Record info for the invoke if requested

// Checks that all direct dependencies of the provided parameters are present in
// the container. Returns an error if not.
func shallowCheckDependencies(c containerStore, pl paramList) error {
	_ = "STUB: not implemented"
	return nil
}

func findMissingDependencies(c containerStore, params ...param) []paramSingle {
	_ = "STUB: not implemented"
	return nil
}

// This means that there is no provider that provides this value,
// and it is NOT being decorated and is NOT optional.
// In the case that there is no providers but there is a decorated value
// of this type, it can be provided safely so we can safely skip this.
