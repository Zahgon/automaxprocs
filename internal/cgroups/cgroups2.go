// Copyright (c) 2022 Uber Technologies, Inc.
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

import (
	"errors"
)

const (
	// _cgroupv2CPUMax is the file name for the CGroup-V2 CPU max and period
	// parameter.
	_cgroupv2CPUMax = "cpu.max"
	// _cgroupFSType is the Linux CGroup-V2 file system type used in
	// `/proc/$PID/mountinfo`.
	_cgroupv2FSType = "cgroup2"

	_cgroupv2MountPoint = "/sys/fs/cgroup"

	_cgroupV2CPUMaxDefaultPeriod = 100000
	_cgroupV2CPUMaxQuotaMax      = "max"
)

const (
	_cgroupv2CPUMaxQuotaIndex = iota
	_cgroupv2CPUMaxPeriodIndex
)

// ErrNotV2 indicates that the system is not using cgroups2.
var ErrNotV2 = errors.New("not using cgroups2")

// CGroups2 provides access to cgroups data for systems using cgroups2.
type CGroups2 struct {
	mountPoint string
	groupPath  string
	cpuMaxFile string
}

// NewCGroups2ForCurrentProcess builds a CGroups2 for the current process.
//
// This returns ErrNotV2 if the system is not using cgroups2.
func NewCGroups2ForCurrentProcess() (*CGroups2, error) { _ = "STUB: not implemented"; return nil, nil }

func newCGroups2From(mountInfoPath, procPathCGroup string) (*CGroups2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find v2 subsystem by looking for the `0` id

func isCGroupV2(procPathMountInfo string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CPUQuota returns the CPU quota applied with the CPU cgroup2 controller.
// It is a result of reading cpu quota and period from cpu.max file.
// It will return `cpu.max / cpu.period`. If cpu.max is set to max, it returns
// (-1, false, nil)
func (cg *CGroups2) CPUQuota() (float64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}
