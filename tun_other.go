//go:build !linux && !windows

package tun

import (
	"net/netip"
	"os"
)

func Open(name string, inet4Address netip.Prefix, inet6Address netip.Prefix, mtu uint32, autoRoute bool) (Tun, error) {
	return nil, os.ErrInvalid
}
