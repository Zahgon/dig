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

// A ProvideOption modifies the default behavior of Provide.
type ProvideOption interface {
	applyProvideOption(*provideOptions)
}

type provideOptions struct {
	Name           string
	Group          string
	Info           *ProvideInfo
	As             []interface{}
	Location       *digreflect.Func
	Exported       bool
	Callback       Callback
	BeforeCallback BeforeCallback
}

func (o *provideOptions) Validate() error { _ = "STUB: not implemented"; return nil }

// Names must be representable inside a backquoted string. The only
// limitation for raw string literals as per
// https://golang.org/ref/spec#raw_string_lit is that they cannot contain
// backquotes.

// Name is a ProvideOption that specifies that all values produced by a
// constructor should have the given name. See also the package documentation
// about Named Values.
//
// Given,
//
//	func NewReadOnlyConnection(...) (*Connection, error)
//	func NewReadWriteConnection(...) (*Connection, error)
//
// The following will provide two connections to the container: one under the
// name "ro" and the other under the name "rw".
//
//	c.Provide(NewReadOnlyConnection, dig.Name("ro"))
//	c.Provide(NewReadWriteConnection, dig.Name("rw"))
//
// This option cannot be provided for constructors which produce result
// objects.
func Name(name string) ProvideOption { _ = "STUB: not implemented"; return *new(ProvideOption) }

type provideNameOption string

func (o provideNameOption) String() string { _ = "STUB: not implemented"; return "" }

func (o provideNameOption) applyProvideOption(opt *provideOptions) {
	_ = "STUB: not implemented"
	return

	// Group is a ProvideOption that specifies that all values produced by a
	// constructor should be added to the specified group. See also the package
	// documentation about Value Groups.
	//
	// This option cannot be provided for constructors which produce result
	// objects.
}

func Group(group string) ProvideOption { _ = "STUB: not implemented"; return *new(ProvideOption) }

type provideGroupOption string

func (o provideGroupOption) String() string { _ = "STUB: not implemented"; return "" }

func (o provideGroupOption) applyProvideOption(opt *provideOptions) {
	_ = "STUB: not implemented"
	return

	// ID is a unique integer representing the constructor node in the dependency graph.
}

type ID int

// ProvideInfo provides information about the constructor's inputs and outputs
// types as strings, as well as the ID of the constructor supplied to the Container.
// It contains ID for the constructor, as well as slices of Input and Output types,
// which are Stringers that report the types of the parameters and results respectively.
type ProvideInfo struct {
	ID      ID
	Inputs  []*Input
	Outputs []*Output
}

// Input contains information on an input parameter of a function.
type Input struct {
	t           reflect.Type
	optional    bool
	name, group string
}

func (i *Input) String() string { _ = "STUB: not implemented"; return "" }

// Output contains information on an output produced by a function.
type Output struct {
	t           reflect.Type
	name, group string
}

func (o *Output) String() string { _ = "STUB: not implemented"; return "" }

// FillProvideInfo is a ProvideOption that writes info on what Dig was able to get
// out of the provided constructor into the provided ProvideInfo.
func FillProvideInfo(info *ProvideInfo) ProvideOption {
	_ = "STUB: not implemented"
	return *new(ProvideOption)
}

type fillProvideInfoOption struct{ info *ProvideInfo }

func (o fillProvideInfoOption) String() string { _ = "STUB: not implemented"; return "" }

func (o fillProvideInfoOption) applyProvideOption(opts *provideOptions) {
	_ = "STUB: not implemented"
	return

	// As is a ProvideOption that specifies that the value produced by the
	// constructor implements one or more other interfaces and is provided
	// to the container as those interfaces.
	//
	// As expects one or more pointers to the implemented interfaces. Values
	// produced by constructors will be then available in the container as
	// implementations of all of those interfaces, but not as the value itself.
	//
	// For example, the following will make io.Reader and io.Writer available
	// in the container, but not buffer.
	//
	//	c.Provide(newBuffer, dig.As(new(io.Reader), new(io.Writer)))
	//
	// That is, the above is equivalent to the following.
	//
	//	c.Provide(func(...) (io.Reader, io.Writer) {
	//	  b := newBuffer(...)
	//	  return b, b
	//	})
	//
	// If used with dig.Name, the type produced by the constructor and the types
	// specified with dig.As will all use the same name. For example,
	//
	//	c.Provide(newFile, dig.As(new(io.Reader)), dig.Name("temp"))
	//
	// The above is equivalent to the following.
	//
	//	type Result struct {
	//	  dig.Out
	//
	//	  Reader io.Reader `name:"temp"`
	//	}
	//
	//	c.Provide(func(...) Result {
	//	  f := newFile(...)
	//	  return Result{
	//	    Reader: f,
	//	  }
	//	})
	//
	// This option cannot be provided for constructors which produce result
	// objects.
}

func As(i ...interface{}) ProvideOption { _ = "STUB: not implemented"; return *new(ProvideOption) }

type provideAsOption []interface{}

func (o provideAsOption) String() string { _ = "STUB: not implemented"; return "" }

func (o provideAsOption) applyProvideOption(opts *provideOptions) {
	_ = "STUB: not implemented"
	return
}

// LocationForPC is a ProvideOption which specifies an alternate function program
// counter address to be used for debug information. The package, name, file and
// line number of this alternate function address will be used in error messages
// and DOT graphs. This option is intended to be used with functions created
// with the reflect.MakeFunc method whose error messages are otherwise hard to
// understand
func LocationForPC(pc uintptr) ProvideOption { _ = "STUB: not implemented"; return *new(ProvideOption) }

type provideLocationOption struct{ loc *digreflect.Func }

func (o provideLocationOption) String() string { _ = "STUB: not implemented"; return "" }

func (o provideLocationOption) applyProvideOption(opts *provideOptions) {
	_ = "STUB: not implemented"
	return

	// Export is a ProvideOption which specifies that the provided function should
	// be made available to all Scopes available in the application, regardless
	// of which Scope it was provided from. By default, it is false.
	//
	// For example,
	//
	//	c := New()
	//	s1 := c.Scope("child 1")
	//	s2:= c.Scope("child 2")
	//	s1.Provide(func() *bytes.Buffer { ... })
	//
	// does not allow the constructor returning *bytes.Buffer to be made available to
	// the root Container c or its sibling Scope s2.
	//
	// With Export, you can make this constructor available to all the Scopes:
	//
	//	s1.Provide(func() *bytes.Buffer { ... }, Export(true))
}

func Export(export bool) ProvideOption { _ = "STUB: not implemented"; return *new(ProvideOption) }

type provideExportOption struct{ exported bool }

func (o provideExportOption) String() string { _ = "STUB: not implemented"; return "" }

func (o provideExportOption) applyProvideOption(opts *provideOptions) {
	_ = "STUB: not implemented"
	return
}

// provider encapsulates a user-provided constructor.
type provider interface {
	// ID is a unique numerical identifier for this provider.
	ID() dot.CtorID

	// Order reports the order of this provider in the graphHolder.
	// This value is usually returned by the graphHolder.NewNode method.
	Order(*Scope) int

	// Location returns where this constructor was defined.
	Location() *digreflect.Func

	// ParamList returns information about the direct dependencies of this
	// constructor.
	ParamList() paramList

	// ResultList returns information about the values produced by this
	// constructor.
	ResultList() resultList

	// Calls the underlying constructor, reading values from the
	// containerStore as needed.
	//
	// The values produced by this provider should be submitted into the
	// containerStore.
	Call(containerStore) error

	CType() reflect.Type

	OrigScope() *Scope
}

// Provide teaches the container how to build values of one or more types and
// expresses their dependencies.
//
// The first argument of Provide is a function that accepts zero or more
// parameters and returns one or more results. The function may optionally
// return an error to indicate that it failed to build the value. This
// function will be treated as the constructor for all the types it returns.
// This function will be called AT MOST ONCE when a type produced by it, or a
// type that consumes this function's output, is requested via Invoke. If the
// same types are requested multiple times, the previously produced value will
// be reused.
//
// Provide accepts argument types or dig.In structs as dependencies, and
// separate return values or dig.Out structs for results.
func (c *Container) Provide(constructor interface{}, opts ...ProvideOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Provide teaches the Scope how to build values of one or more types and
// expresses their dependencies.
//
// The first argument of Provide is a function that accepts zero or more
// parameters and returns one or more results. The function may optionally
// return an error to indicate that it failed to build the value. This
// function will be treated as the constructor for all the types it returns.
// This function will be called AT MOST ONCE when a type produced by it, or a
// type that consumes this function's output, is requested via Invoke. If the
// same types are requested multiple times, the previously produced value will
// be reused.
//
// Provide accepts argument types or dig.In structs as dependencies, and
// separate return values or dig.Out structs for results.
//
// When a constructor is Provided to a Scope, it will propagate this to any
// Scopes that are descendents, but not ancestors of this Scope.
// To provide a constructor to all the Scopes available, provide it to
// Container, which is the root Scope.
func (s *Scope) Provide(constructor interface{}, opts ...ProvideOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scope) provide(ctor interface{}, opts provideOptions) (err error) {
	_ = "STUB: not implemented"
	// If Export option is provided to the constructor, this should be injected to the
	// root-level Scope (Container) to allow it to propagate to all other Scopes.
	return nil
}

// For all scopes affected by this change,
// take a snapshot of the current graph state before
// we start making changes to it as we may need to
// undo them upon encountering errors.

// Cache old providers before running cycle detection.

// When a cycle is detected, recover the old providers to reset
// the providers map back to what it was before this node was
// introduced.

// Record introspection info for caller if Info option is specified

// Builds a collection of all result types produced by this constructor.
func (s *Scope) findAndValidateResults(rl resultList) (map[key]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Visits the results of a node and compiles a collection of all the keys
// produced by that node.
type connectionVisitor struct {
	s *Scope

	// If this points to a non-nil value, we've already encountered an error
	// and should stop traversing.
	err *error

	// Map of keys provided to path that provided this. The path is a string
	// documenting which positional return value or dig.Out attribute is
	// providing this particular key.
	//
	// For example, "[0].Foo" indicates that the value was provided by the Foo
	// attribute of the dig.Out returned as the first result of the
	// constructor.
	keyPaths map[key]string

	// We track the path to the current result here. For example, this will
	// be, ["[1]", "Foo", "Bar"] when we're visiting Bar in,
	//
	//   func() (io.Writer, struct {
	//     dig.Out
	//
	//     Foo struct {
	//       dig.Out
	//
	//       Bar io.Reader
	//     }
	//   })
	currentResultPath []string
}

func (cv connectionVisitor) AnnotateWithField(f resultObjectField) resultVisitor {
	_ = "STUB: not implemented"
	return *new(resultVisitor)
}

func (cv connectionVisitor) AnnotateWithPosition(i int) resultVisitor {
	_ = "STUB: not implemented"
	return *new(resultVisitor)
}

func (cv connectionVisitor) Visit(res result) resultVisitor {
	_ = "STUB: not implemented"
	// Already failed. Stop looking.
	return *new(resultVisitor)
}

// we don't really care about the path for this since conflicts are
// okay for group results. We'll track it for the sake of having a
// value there.

func (cv connectionVisitor) checkKey(k key, path string) error {
	_ = "STUB: not implemented"
	return nil
}
