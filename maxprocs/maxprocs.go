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

// Package maxprocs lets Go programs easily configure runtime.GOMAXPROCS to
// match the configured Linux CPU quota. Unlike the top-level automaxprocs
// package, it lets the caller configure logging and handle errors.
package maxprocs // import "go.uber.org/automaxprocs/maxprocs"

import (
	iruntime "go.uber.org/automaxprocs/internal/runtime"
)

const _maxProcsKey = "GOMAXPROCS"

func currentMaxProcs() int { _ = "STUB: not implemented"; return 0 }

type config struct {
	printf         func(string, ...interface{})
	procs          func(int, func(v float64) int) (int, iruntime.CPUQuotaStatus, error)
	minGOMAXPROCS  int
	roundQuotaFunc func(v float64) int
}

func (c *config) log(fmt string, args ...interface{}) { _ = "STUB: not implemented"; return }

// An Option alters the behavior of Set.
type Option interface {
	apply(*config)
}

// Logger uses the supplied printf implementation for log output. By default,
// Set doesn't log anything.
func Logger(printf func(string, ...interface{})) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Min sets the minimum GOMAXPROCS value that will be used.
// Any value below 1 is ignored.
func Min(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

// RoundQuotaFunc sets the function that will be used to covert the CPU quota from float to int.
func RoundQuotaFunc(rf func(v float64) int) Option { _ = "STUB: not implemented"; return *new(Option) }

type optionFunc func(*config)

func (of optionFunc) apply(cfg *config) {
	_ = "STUB: not implemented"

	// Set GOMAXPROCS to match the Linux container CPU quota (if any), returning
	// any error encountered and an undo function.
	//
	// Set is a no-op on non-Linux systems and in Linux environments without a
	// configured CPU quota.
	return
}

func Set(opts ...Option) (func(), error) { _ = "STUB: not implemented"; return nil, nil }

// Honor the GOMAXPROCS environment variable if present. Otherwise, amend
// `runtime.GOMAXPROCS()` with the current process' CPU quota if the OS is
// Linux, and guarantee a minimum value of 1. The minimum guaranteed value
// can be overridden using `maxprocs.Min()`.
