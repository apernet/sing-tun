//go:build linux

package tun

import "golang.org/x/sys/unix"

var errBadFd = unix.EBADFD
