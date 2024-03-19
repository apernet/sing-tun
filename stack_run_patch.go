package tun

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
	return s.tunLoop()
}
