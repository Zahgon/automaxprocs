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

const (
	// _cgroupFSType is the Linux CGroup file system type used in
	// `/proc/$PID/mountinfo`.
	_cgroupFSType = "cgroup"
	// _cgroupSubsysCPU is the CPU CGroup subsystem.
	_cgroupSubsysCPU = "cpu"
	// _cgroupSubsysCPUAcct is the CPU accounting CGroup subsystem.
	_cgroupSubsysCPUAcct = "cpuacct"
	// _cgroupSubsysCPUSet is the CPUSet CGroup subsystem.
	_cgroupSubsysCPUSet = "cpuset"
	// _cgroupSubsysMemory is the Memory CGroup subsystem.
	_cgroupSubsysMemory = "memory"

	// _cgroupCPUCFSQuotaUsParam is the file name for the CGroup CFS quota
	// parameter.
	_cgroupCPUCFSQuotaUsParam = "cpu.cfs_quota_us"
	// _cgroupCPUCFSPeriodUsParam is the file name for the CGroup CFS period
	// parameter.
	_cgroupCPUCFSPeriodUsParam = "cpu.cfs_period_us"
)

const (
	_procPathCGroup    = "/proc/self/cgroup"
	_procPathMountInfo = "/proc/self/mountinfo"
)

// CGroups is a map that associates each CGroup with its subsystem name.
type CGroups map[string]*CGroup

// NewCGroups returns a new *CGroups from given `mountinfo` and `cgroup` files
// under for some process under `/proc` file system (see also proc(5) for more
// information).
func NewCGroups(procPathMountInfo, procPathCGroup string) (CGroups, error) {
	_ = "STUB: not implemented"
	return *new(CGroups), nil
}

// NewCGroupsForCurrentProcess returns a new *CGroups instance for the current
// process.
func NewCGroupsForCurrentProcess() (CGroups, error) {
	_ = "STUB: not implemented"
	return *new(CGroups), nil
}

// CPUQuota returns the CPU quota applied with the CPU cgroup controller.
// It is a result of `cpu.cfs_quota_us / cpu.cfs_period_us`. If the value of
// `cpu.cfs_quota_us` was not set (-1), the method returns `(-1, nil)`.
func (cg CGroups) CPUQuota() (float64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}
