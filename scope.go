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
	"math/rand"
	"reflect"

	"go.uber.org/dig/internal/digclock"
)

// A ScopeOption modifies the default behavior of Scope; currently,
// there are no implementations.
type ScopeOption interface {
	noScopeOption() // yet
}

// Scope is a scoped DAG of types and their dependencies.
// A Scope may also have one or more child Scopes that inherit
// from it.
type Scope struct {
	// This implements containerStore interface.

	// Name of the Scope
	name string
	// Mapping from key to all the constructor node that can provide a value for that
	// key.
	providers map[key][]*constructorNode

	// Mapping from key to the decorator that decorates a value for that key.
	decorators map[key]*decoratorNode

	// constructorNodes provided directly to this Scope. i.e. it does not include
	// any nodes that were provided to the parent Scope this inherited from.
	nodes []*constructorNode

	// Values that generated via decorators in the Scope.
	decoratedValues map[key]reflect.Value

	// Values that generated directly in the Scope.
	values map[key]reflect.Value

	// Values groups that generated directly in the Scope.
	groups map[key][]reflect.Value

	// Values groups that generated via decoraters in the Scope.
	decoratedGroups map[key]reflect.Value

	// Source of randomness.
	rand *rand.Rand

	// Flag indicating whether the graph has been checked for cycles.
	isVerifiedAcyclic bool

	// Defer acyclic check on provide until Invoke.
	deferAcyclicVerification bool

	// Recover from panics in user-provided code and wrap in an exported error type.
	recoverFromPanics bool

	// invokerFn calls a function with arguments provided to Provide or Invoke.
	invokerFn invokerFn

	// graph of this Scope. Note that this holds the dependency graph of all the
	// nodes that affect this Scope, not just the ones provided directly to this Scope.
	gh *graphHolder

	// Parent of this Scope.
	parentScope *Scope

	// All the child scopes of this Scope.
	childScopes []*Scope

	// clockSrc stores the source of time. Defaults to system clock.
	clockSrc digclock.Clock
}

func newScope() *Scope { _ = "STUB: not implemented"; return nil }

// Scope creates a new Scope with the given name and options from current Scope.
// Any constructors that the current Scope knows about, as well as any modifications
// made to it in the future will be propagated to the child scope.
// However, no modifications made to the child scope being created will be propagated
// to the parent Scope.
func (s *Scope) Scope(name string, opts ...ScopeOption) *Scope {
	_ = "STUB: not implemented"
	return nil
}

// child copies the parent's graph nodes.

// ancestors returns a list of scopes of ancestors of this scope up to the
// root. The scope at at index 0 is this scope itself.
func (s *Scope) ancestors() []*Scope { _ = "STUB: not implemented"; return nil }

func (s *Scope) appendSubscopes(dest []*Scope) []*Scope { _ = "STUB: not implemented"; return nil }

func (s *Scope) storesToRoot() []containerStore { _ = "STUB: not implemented"; return nil }

func (s *Scope) knownTypes() []reflect.Type { _ = "STUB: not implemented"; return nil }

func (s *Scope) getValue(name string, t reflect.Type) (v reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func (s *Scope) getDecoratedValue(name string, t reflect.Type) (v reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func (s *Scope) setValue(name string, t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (s *Scope) setDecoratedValue(name string, t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (s *Scope) getValueGroup(name string, t reflect.Type) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

// shuffle the list so users don't rely on the ordering of grouped values

func (s *Scope) getDecoratedValueGroup(name string, t reflect.Type) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func (s *Scope) submitGroupedValue(name string, t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (s *Scope) submitDecoratedGroupedValue(name string, t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (s *Scope) getValueProviders(name string, t reflect.Type) []provider {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scope) getGroupProviders(name string, t reflect.Type) []provider {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scope) getValueDecorator(name string, t reflect.Type) (decorator, bool) {
	_ = "STUB: not implemented"
	return *new(decorator), false
}

func (s *Scope) getGroupDecorator(name string, t reflect.Type) (decorator, bool) {
	_ = "STUB: not implemented"
	return *new(decorator), false
}

func (s *Scope) getDecorators(k key) (decorator, bool) {
	_ = "STUB: not implemented"
	return *new(decorator), false
}

func (s *Scope) getProviders(k key) []provider { _ = "STUB: not implemented"; return nil }

func (s *Scope) getAllGroupProviders(name string, t reflect.Type) []provider {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scope) getAllValueProviders(name string, t reflect.Type) []provider {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scope) getAllProviders(k key) []provider { _ = "STUB: not implemented"; return nil }

func (s *Scope) invoker() invokerFn { _ = "STUB: not implemented"; return *new(invokerFn) }

func (s *Scope) clock() digclock.Clock {
	_ = "STUB: not implemented"

	// adds a new graphNode to this Scope and all of its descendent
	// scope.
	return *new(digclock.Clock)
}

func (s *Scope) newGraphNode(wrapped interface{}, orders map[*Scope]int) {
	_ = "STUB: not implemented"
	return
}

func (s *Scope) cycleDetectedError(cycle []int) error { _ = "STUB: not implemented"; return nil }

// Returns the root Scope that can be reached from this Scope.
func (s *Scope) rootScope() *Scope { _ = "STUB: not implemented"; return nil }

// String representation of the entire Scope
func (s *Scope) String() string { _ = "STUB: not implemented"; return "" }
