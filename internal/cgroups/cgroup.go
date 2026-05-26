// Copyright (c) 2017 Uber Technologies, Inc.
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

//go:build linux
// +build linux

package cgroups

// CGroup represents the data structure for a Linux control group.
type CGroup struct {
	path string
}

// NewCGroup returns a new *CGroup from a given path.
func NewCGroup(path string) *CGroup { _ = "STUB: not implemented"; return nil }

// Path returns the path of the CGroup*.
func (cg *CGroup) Path() string {
	_ = "STUB: not implemented"

	// ParamPath returns the path of the given cgroup param under itself.
	return ""
}

func (cg *CGroup) ParamPath(param string) string { _ = "STUB: not implemented"; return "" }

// readFirstLine reads the first line from a cgroup param file.
func (cg *CGroup) readFirstLine(param string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// readInt parses the first line from a cgroup param file as int.
func (cg *CGroup) readInt(param string) (int, error) { _ = "STUB: not implemented"; return 0, nil }
