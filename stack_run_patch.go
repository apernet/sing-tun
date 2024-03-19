package tun

import (
	"errors"
)

type StackRunner interface {
	Stack
	Run() error
}

var _ StackRunner = (*System)(nil)

func (s *System) Run() error {
	err := s.Start()
	defer s.Close()
	if err != nil {
		return err
	}
	s.tunLoop()
	return errors.New("tun loop exited")
}
