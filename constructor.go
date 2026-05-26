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

import (
	"reflect"

	"go.uber.org/dig/internal/digreflect"
	"go.uber.org/dig/internal/dot"
)

// constructorNode is a node in the dependency graph that represents
// a constructor provided by the user.
//
// constructorNodes can produce zero or more values that they store into the container.
// For the Provide path, we verify that constructorNodes produce at least one value,
// otherwise the function will never be called.
type constructorNode struct {
	ctor  interface{}
	ctype reflect.Type

	// Location where this function was defined.
	location *digreflect.Func

	// id uniquely identifies the constructor that produces a node.
	id dot.CtorID

	// Whether the constructor owned by this node was already called.
	called bool

	// Type information about constructor parameters.
	paramList paramList

	// Type information about constructor results.
	resultList resultList

	// Order of this node in each Scopes' graphHolders.
	orders map[*Scope]int

	// Scope this node is part of.
	s *Scope

	// Scope this node was originally provided to.
	// This is different from s if and only if the constructor was Provided with ExportOption.
	origS *Scope

	// Callback for this provided function, if there is one.
	callback Callback

	// BeforeCallback for this provided function, if there is one.
	beforeCallback BeforeCallback
}

type constructorOptions struct {
	// If specified, all values produced by this constructor have the provided name
	// belong to the specified value group or implement any of the interfaces.
	ResultName     string
	ResultGroup    string
	ResultAs       []interface{}
	Location       *digreflect.Func
	Callback       Callback
	BeforeCallback BeforeCallback
}

func newConstructorNode(ctor interface{}, s *Scope, origS *Scope, opts constructorOptions) (*constructorNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *constructorNode) Location() *digreflect.Func { _ = "STUB: not implemented"; return nil }
func (n *constructorNode) ParamList() paramList       { _ = "STUB: not implemented"; return *new(paramList) }
func (n *constructorNode) ResultList() resultList {
	_ = "STUB: not implemented"
	return *new(resultList)
}
func (n *constructorNode) ID() dot.CtorID { _ = "STUB: not implemented"; return *new(dot.CtorID) }
func (n *constructorNode) CType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}
func (n *constructorNode) Order(s *Scope) int { _ = "STUB: not implemented"; return 0 }
func (n *constructorNode) OrigScope() *Scope  { _ = "STUB: not implemented"; return nil }

// CopyOrder copies the order for the given parent scope to the given child scope.
func (n *constructorNode) CopyOrder(parent, child *Scope) { _ = "STUB: not implemented"; return }

func (n *constructorNode) String() string { _ = "STUB: not implemented"; return "" }

// Call calls this constructor if it hasn't already been called and
// injects any values produced by it into the provided container.
func (n *constructorNode) Call(c containerStore) (err error) { _ = "STUB: not implemented"; return nil }

// Wrap in separate func to include PanicErrors

/* decorating */

// Commit the result to the original container that this constructor
// was supplied to. The provided constructor is only used for a view of
// the rest of the graph to instantiate the dependencies of this
// container.

// stagingContainerWriter is a containerWriter that records the changes that
// would be made to a containerWriter and defers them until Commit is called.
type stagingContainerWriter struct {
	values map[key]reflect.Value
	groups map[key][]reflect.Value
}

var _ containerWriter = (*stagingContainerWriter)(nil)

func newStagingContainerWriter() *stagingContainerWriter { _ = "STUB: not implemented"; return nil }

func (sr *stagingContainerWriter) setValue(name string, t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (sr *stagingContainerWriter) setDecoratedValue(_ string, _ reflect.Type, _ reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (sr *stagingContainerWriter) submitGroupedValue(group string, t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (sr *stagingContainerWriter) submitDecoratedGroupedValue(_ string, _ reflect.Type, _ reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// Commit commits the received results to the provided containerWriter.
func (sr *stagingContainerWriter) Commit(cw containerWriter) { _ = "STUB: not implemented"; return }
