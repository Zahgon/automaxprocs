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
	_mountInfoSep               = " "
	_mountInfoOptsSep           = ","
	_mountInfoOptionalFieldsSep = "-"
)

const (
	_miFieldIDMountID = iota
	_miFieldIDParentID
	_miFieldIDDeviceID
	_miFieldIDRoot
	_miFieldIDMountPoint
	_miFieldIDOptions
	_miFieldIDOptionalFields

	_miFieldCountFirstHalf
)

const (
	_miFieldOffsetFSType = iota
	_miFieldOffsetMountSource
	_miFieldOffsetSuperOptions

	_miFieldCountSecondHalf
)

const _miFieldCountMin = _miFieldCountFirstHalf + _miFieldCountSecondHalf

// MountPoint is the data structure for the mount points in
// `/proc/$PID/mountinfo`. See also proc(5) for more information.
type MountPoint struct {
	MountID        int
	ParentID       int
	DeviceID       string
	Root           string
	MountPoint     string
	Options        []string
	OptionalFields []string
	FSType         string
	MountSource    string
	SuperOptions   []string
}

// NewMountPointFromLine parses a line read from `/proc/$PID/mountinfo` and
// returns a new *MountPoint.
func NewMountPointFromLine(line string) (*MountPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// End of optional fields.

// Now we know where the optional fields end, split the line again with a
// limit to avoid issues with spaces in super options as present on WSL.

// Translate converts an absolute path inside the *MountPoint's file system to
// the host file system path in the mount namespace the *MountPoint belongs to.
func (mp *MountPoint) Translate(absPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// parseMountInfo parses procPathMountInfo (usually at `/proc/$PID/mountinfo`)
// and yields parsed *MountPoint into newMountPoint.
func parseMountInfo(procPathMountInfo string, newMountPoint func(*MountPoint) error) error {
	_ = "STUB: not implemented"
	return nil
}
