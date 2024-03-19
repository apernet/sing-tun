//go:build !linux

package tun

import "errors"

var errBadFd = errors.New("EBADFD placeholder (not exists on this platform)")
