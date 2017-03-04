//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// This file is machine generated by the command:
//   mksyscall.pl serial_darwin.go
// The generated stub is modified to make it compile under the "serial" package

package serial // import "go.bug.st/serial.v1"

import "syscall"

func ioctl(fd int, req uint64, data uintptr) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(data))
	if e1 != 0 {
		err = e1
	}
	return
}
