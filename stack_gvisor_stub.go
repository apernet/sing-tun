
package tun

import E "github.com/sagernet/sing/common/exceptions"

const WithGVisor = false

var ErrGVisorNotIncluded = E.New(`gVisor is not supported in this fork.`)

func NewGVisor(
	options StackOptions,
) (Stack, error) {
	return nil, ErrGVisorNotIncluded
}

func NewMixed(
	options StackOptions,
) (Stack, error) {
	return nil, ErrGVisorNotIncluded
}
