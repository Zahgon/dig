// Copyright (c) 2019-2021 Uber Technologies, Inc.
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
	"reflect"

	"go.uber.org/dig/internal/dot"
)

// The param interface represents a dependency for a constructor.
//
// The following implementations exist:
//
//	paramList     All arguments of the constructor.
//	paramSingle   An explicitly requested type.
//	paramObject   dig.In struct where each field in the struct can be another
//	              param.
//	paramGroupedSlice
//	              A slice consuming a value group. This will receive all
//	              values produced with a `group:".."` tag with the same name
//	              as a slice.
type param interface {
	fmt.Stringer

	// Build this dependency and any of its dependencies from the provided
	// Container.
	//
	// This MAY panic if the param does not produce a single value.
	Build(store containerStore) (reflect.Value, error)

	// DotParam returns a slice of dot.Param(s).
	DotParam() []*dot.Param
}

var (
	_ param = paramSingle{}
	_ param = paramObject{}
	_ param = paramList{}
	_ param = paramGroupedSlice{}
)

// newParam builds a param from the given type. If the provided type is a
// dig.In struct, an paramObject will be returned.
func newParam(t reflect.Type, c containerStore) (param, error) {
	_ = "STUB: not implemented"
	return *new(param), nil
}

// paramList holds all arguments of the constructor as params.
//
// NOTE: Build() MUST NOT be called on paramList. Instead, BuildList
// must be called.
type paramList struct {
	ctype reflect.Type // type of the constructor

	Params []param
}

func (pl paramList) DotParam() []*dot.Param { _ = "STUB: not implemented"; return nil }

func (pl paramList) String() string { _ = "STUB: not implemented"; return "" }

// newParamList builds a paramList from the provided constructor type.
//
// Variadic arguments of a constructor are ignored and not included as
// dependencies.
func newParamList(ctype reflect.Type, c containerStore) (paramList, error) {
	_ = "STUB: not implemented"
	return *new(paramList), nil
}

// NOTE: If the function is variadic, we skip the last argument
// because we're not filling variadic arguments yet. See #120.

func (pl paramList) Build(containerStore) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Unreachable, as BugPanicf above will panic.

// BuildList returns an ordered list of values which may be passed directly
// to the underlying constructor.
func (pl paramList) BuildList(c containerStore) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// paramSingle is an explicitly requested type, optionally with a name.
//
// This object must be present in the graph as-is unless it's specified as
// optional.
type paramSingle struct {
	Name     string
	Optional bool
	Type     reflect.Type
}

func (ps paramSingle) DotParam() []*dot.Param { _ = "STUB: not implemented"; return nil }

func (ps paramSingle) String() string {
	_ = "STUB: not implemented"
	// tally.Scope[optional] means optional
	// tally.Scope[optional, name="foo"] means named optional
	return ""
}

// search the given container and its ancestors for a decorated value.
func (ps paramSingle) getDecoratedValue(c containerStore) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// builds the parameter using decorators in all scopes that affect the
// current scope, if there are any. If there are multiple Scopes that decorates
// this parameter, the closest one to the Scope that invoked this will be used.
// If there are no decorators associated with this parameter, _noValue is returned.
func (ps paramSingle) buildWithDecorators(c containerStore) (v reflect.Value, found bool, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

// This decorator is already being run.
// Avoid a cycle and look further.

func (ps paramSingle) Build(c containerStore) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Check whether the value is a decorated value first.

// Starting at the given container and working our way up its parents,
// find one that provides this dependency.
//
// Once found, we'll use that container for the rest of the invocation.
// Dependencies of this type will begin searching at that container,
// rather than starting at base.

// first check if the scope already has cached a value for the type.

// If we're missing dependencies but the parameter itself is optional,
// we can just move on.

// If we get here, it's impossible for the value to be absent from the
// container.

// paramObject is a dig.In struct where each field is another param.
//
// This object is not expected in the graph as-is.
type paramObject struct {
	Type        reflect.Type
	Fields      []paramObjectField
	FieldOrders []int
}

func (po paramObject) DotParam() []*dot.Param { _ = "STUB: not implemented"; return nil }

func (po paramObject) String() string { _ = "STUB: not implemented"; return "" }

// getParamOrder returns the order(s) of a parameter type.
func getParamOrder(gh *graphHolder, param param) []int { _ = "STUB: not implemented"; return nil }

// value group parameters have nodes of their own.
// We can directly return that here.

// newParamObject builds an paramObject from the provided type. The type MUST
// be a dig.In struct.
func newParamObject(t reflect.Type, c containerStore) (paramObject, error) {
	_ = "STUB: not implemented"
	return *new(paramObject), nil
}

// Check if the In type supports ignoring unexported fields.

// Skip over the dig.In embed.

// Skip over an unexported field if it is allowed.

func (po paramObject) Build(c containerStore) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// We have to build soft groups after all other fields, to avoid cases
// when a field calls a provider for a soft value group, but the value is
// not provided to it because the value group is declared before the field

// paramObjectField is a single field of a dig.In struct.
type paramObjectField struct {
	// Name of the field in the struct.
	FieldName string

	// Index of this field in the target struct.
	//
	// We need to track this separately because not all fields of the
	// struct map to params.
	FieldIndex int

	// The dependency requested by this field.
	Param param
}

func (pof paramObjectField) DotParam() []*dot.Param { _ = "STUB: not implemented"; return nil }

func newParamObjectField(idx int, f reflect.StructField, c containerStore) (paramObjectField, error) {
	_ = "STUB: not implemented"
	return *new(paramObjectField), nil
}

func (pof paramObjectField) Build(c containerStore) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// paramGroupedSlice is a param which produces a slice of values with the same
// group name.
type paramGroupedSlice struct {
	// Name of the group as specified in the `group:".."` tag.
	Group string

	// Type of the slice.
	Type reflect.Type

	// Soft is used to denote a soft dependency between this param and its
	// constructors, if it's true its constructors are only called if they
	// provide another value requested in the graph
	Soft bool

	orders map[*Scope]int
}

func (pt paramGroupedSlice) String() string {
	_ = "STUB: not implemented"
	// io.Reader[group="foo"] refers to a group of io.Readers called 'foo'
	return ""
}

func (pt paramGroupedSlice) DotParam() []*dot.Param { _ = "STUB: not implemented"; return nil }

// newParamGroupedSlice builds a paramGroupedSlice from the provided type with
// the given name.
//
// The type MUST be a slice type.
func newParamGroupedSlice(f reflect.StructField, c containerStore) (paramGroupedSlice, error) {
	_ = "STUB: not implemented"
	return *new(paramGroupedSlice), nil
}

// retrieves any decorated values that may be committed in this scope, or
// any of the parent Scopes. In the case where there are multiple scopes that
// are decorating the same type, the closest scope in effect will be replacing
// any decorated value groups provided in further scopes.
func (pt paramGroupedSlice) getDecoratedValues(c containerStore) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// search the given container and its parents for matching group decorators
// and call them to commit values. If any decorators return an error,
// that error is returned immediately. If all decorators succeeds, nil is returned.
// The order in which the decorators are invoked is from the top level scope to
// the current scope, to account for decorators that decorate values that were
// already decorated.
func (pt paramGroupedSlice) callGroupDecorators(c containerStore) error {
	_ = "STUB: not implemented"
	return nil
}

// This decorator is already being run. Avoid cycle
// and look further.

// search the given container and its parent for matching group providers and
// call them to commit values. If an error is encountered, return the number
// of providers called and a non-nil error from the first provided.
func (pt paramGroupedSlice) callGroupProviders(c containerStore) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pt paramGroupedSlice) Build(c containerStore) (reflect.Value, error) {
	_ = "STUB: not implemented"
	// do not call this if we are already inside a decorator since
	// it will result in an infinite recursion. (i.e. decorate -> params.BuildList() -> Decorate -> params.BuildList...)
	// this is safe since a value can be decorated at most once in a given scope.
	return *new(reflect.Value), nil
}

// Check if we have decorated values

// If we do not have any decorated values and the group isn't soft,
// find the providers and call them.

// Checks if ignoring unexported files in an In struct is allowed.
// The struct field MUST be an _inType.
func isIgnoreUnexportedSet(f reflect.StructField) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
