package steamcmd

var Quit = &quit{}

type quit struct{}

func (q *quit) String() string {
	return "quit"
}

func (*quit) Check(flags *Flags) error {
	return nil
}

func (*quit) Args() ([]string, error) {
	return []string{"quit"}, nil
}

func (*quit) Modify(_ *Flags) error {
	return nil
}
