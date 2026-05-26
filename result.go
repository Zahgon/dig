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
	"reflect"

	"go.uber.org/dig/internal/dot"
)

// The result interface represents a result produced by a constructor.
//
// The following implementations exist:
//   resultList    All values returned by the constructor.
//   resultSingle  A single value produced by a constructor.
//   resultObject  dig.Out struct where each field in the struct can be
//                 another result.
//   resultGrouped A value produced by a constructor that is part of a value
//                 group.

type result interface {
	// Extracts the values for this result from the provided value and
	// stores them into the provided containerWriter.
	//
	// This MAY panic if the result does not consume a single value.
	Extract(containerWriter, bool, reflect.Value)

	// DotResult returns a slice of dot.Result(s).
	DotResult() []*dot.Result
}

var (
	_ result = resultSingle{}
	_ result = resultObject{}
	_ result = resultList{}
	_ result = resultGrouped{}
)

type resultOptions struct {
	// If set, this is the name of the associated result value.
	//
	// For Result Objects, name:".." tags on fields override this.
	Name  string
	Group string
	As    []interface{}
}

// newResult builds a result from the given type.
func newResult(t reflect.Type, opts resultOptions) (result, error) {
	_ = "STUB: not implemented"
	return *new(result), nil
}

// resultVisitor visits every result in a result tree, allowing tracking state
// at each level.
type resultVisitor interface {
	// Visit is called on the result being visited.
	//
	// If Visit returns a non-nil resultVisitor, that resultVisitor visits all
	// the child results of this result.
	Visit(result) resultVisitor

	// AnnotateWithField is called on each field of a resultObject after
	// visiting it but before walking its descendants.
	//
	// The same resultVisitor is used for all fields: the one returned upon
	// visiting the resultObject.
	//
	// For each visited field, if AnnotateWithField returns a non-nil
	// resultVisitor, it will be used to walk the result of that field.
	AnnotateWithField(resultObjectField) resultVisitor

	// AnnotateWithPosition is called with the index of each result of a
	// resultList after vising it but before walking its descendants.
	//
	// The same resultVisitor is used for all results: the one returned upon
	// visiting the resultList.
	//
	// For each position, if AnnotateWithPosition returns a non-nil
	// resultVisitor, it will be used to walk the result at that index.
	AnnotateWithPosition(idx int) resultVisitor
}

// walkResult walks the result tree for the given result with the provided
// visitor.
//
// resultVisitor.Visit will be called on the provided result and if a non-nil
// resultVisitor is received, it will be used to walk its descendants. If a
// resultObject or resultList was visited, AnnotateWithField and
// AnnotateWithPosition respectively will be called before visiting the
// descendants of that resultObject/resultList.
//
// This is very similar to how go/ast.Walk works.
func walkResult(r result, v resultVisitor) { _ = "STUB: not implemented"; return }

// No sub-results

// resultList holds all values returned by the constructor as results.
type resultList struct {
	ctype reflect.Type

	Results []result

	// For each item at index i returned by the constructor, resultIndexes[i]
	// is the index in .Results for the corresponding result object.
	// resultIndexes[i] is -1 for errors returned by constructors.
	resultIndexes []int
}

func (rl resultList) DotResult() []*dot.Result { _ = "STUB: not implemented"; return nil }

func newResultList(ctype reflect.Type, opts resultOptions) (resultList, error) {
	_ = "STUB: not implemented"
	return *new(resultList), nil
}

func (resultList) Extract(containerWriter, bool, reflect.Value) { _ = "STUB: not implemented"; return }

func (rl resultList) ExtractList(cw containerWriter, decorated bool, values []reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// resultSingle is an explicit value produced by a constructor, optionally
// with a name.
//
// This object will be added to the graph as-is.
type resultSingle struct {
	Name string
	Type reflect.Type

	// If specified, this is a list of types which the value will be made
	// available as, in addition to its own type.
	As []reflect.Type
}

func newResultSingle(t reflect.Type, opts resultOptions) (resultSingle, error) {
	_ = "STUB: not implemented"
	return *new(resultSingle), nil
}

// Special case:
//   c.Provide(func() io.Reader, As(new(io.Reader)))
// Ignore instead of erroring out.

func (rs resultSingle) DotResult() []*dot.Result { _ = "STUB: not implemented"; return nil }

func (rs resultSingle) Extract(cw containerWriter, decorated bool, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// resultObject is a dig.Out struct where each field is another result.
//
// This object is not added to the graph. Its fields are interpreted as
// results and added to the graph if needed.
type resultObject struct {
	Type   reflect.Type
	Fields []resultObjectField
}

func (ro resultObject) DotResult() []*dot.Result { _ = "STUB: not implemented"; return nil }

func newResultObject(t reflect.Type, opts resultOptions) (resultObject, error) {
	_ = "STUB: not implemented"
	return *new(resultObject), nil
}

// Skip over the dig.Out embed.

func (ro resultObject) Extract(cw containerWriter, decorated bool, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// resultObjectField is a single field inside a dig.Out struct.
type resultObjectField struct {
	// Name of the field in the struct.
	FieldName string

	// Index of the field in the struct.
	//
	// We need to track this separately because not all fields of the struct
	// map to results.
	FieldIndex int

	// Result produced by this field.
	Result result
}

func (rof resultObjectField) DotResult() []*dot.Result { _ = "STUB: not implemented"; return nil }

// newResultObjectField(i, f, opts) builds a resultObjectField from the field
// f at index i.
func newResultObjectField(idx int, f reflect.StructField, opts resultOptions) (resultObjectField, error) {
	_ = "STUB: not implemented"
	return *new(resultObjectField), nil
}

// can modify in-place because options are passed-by-value.

// resultGrouped is a value produced by a constructor that is part of a result
// group.
//
// These will be produced as fields of a dig.Out struct.
type resultGrouped struct {
	// Name of the group as specified in the `group:".."` tag.
	Group string

	// Type of value produced.
	Type reflect.Type

	// Indicates elements of a value are to be injected individually, instead of
	// as a group. Requires the value's slice to be a group. If set, Type will be
	// the type of individual elements rather than the group.
	Flatten bool

	// If specified, this is a list of types which the value will be made
	// available as, in addition to its own type.
	As []reflect.Type
}

func (rt resultGrouped) DotResult() []*dot.Result { _ = "STUB: not implemented"; return nil }

// newResultGrouped(f) builds a new resultGrouped from the provided field.
func newResultGrouped(f reflect.StructField) (resultGrouped, error) {
	_ = "STUB: not implemented"
	return *new(resultGrouped), nil
}

func (rt resultGrouped) Extract(cw containerWriter, decorated bool, v reflect.Value) {
	_ = "STUB: not implemented"
	// Decorated values are always flattened.
	return
}
