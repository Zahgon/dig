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
	"io"

	"go.uber.org/dig/internal/dot"
)

// A VisualizeOption modifies the default behavior of Visualize.
type VisualizeOption interface {
	applyVisualizeOption(*visualizeOptions)
}

type visualizeOptions struct {
	VisualizeError error
}

// VisualizeError includes a visualization of the given error in the output of
// Visualize if an error was returned by Invoke or Provide.
//
//	if err := c.Provide(...); err != nil {
//	  dig.Visualize(c, w, dig.VisualizeError(err))
//	}
//
// This option has no effect if the error was nil or if it didn't contain any
// information to visualize.
func VisualizeError(err error) VisualizeOption {
	_ = "STUB: not implemented"
	return *new(VisualizeOption)
}

type visualizeErrorOption struct{ err error }

func (o visualizeErrorOption) String() string { _ = "STUB: not implemented"; return "" }

func (o visualizeErrorOption) applyVisualizeOption(opt *visualizeOptions) {
	_ = "STUB: not implemented"
	return
}

func updateGraph(dg *dot.Graph, err error) error { _ = "STUB: not implemented"; return nil }

// Unwrap error to find the root cause.

// If there are no errVisualizers included, we do not modify the graph.

// We iterate in reverse because the last element is the root cause.

// Remove non-error entries from the graph for readability.

// Visualize parses the graph in Container c into DOT format and writes it to
// io.Writer w.
func Visualize(c *Container, w io.Writer, opts ...VisualizeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func visualizeGraph(w io.Writer, dg *dot.Graph) { _ = "STUB: not implemented"; return }

func visualizeGroup(w io.Writer, g *dot.Group) { _ = "STUB: not implemented"; return }

func visualizeCtor(w io.Writer, index int, c *dot.Ctor) { _ = "STUB: not implemented"; return }

// CanVisualizeError returns true if the error is an errVisualizer.
func CanVisualizeError(err error) bool { _ = "STUB: not implemented"; return false }

func (c *Container) createGraph() *dot.Graph { _ = "STUB: not implemented"; return nil }

func (s *Scope) createGraph() *dot.Graph { _ = "STUB: not implemented"; return nil }

func (s *Scope) addNodes(dg *dot.Graph) { _ = "STUB: not implemented"; return }

func newDotCtor(n *constructorNode) *dot.Ctor { _ = "STUB: not implemented"; return nil }
