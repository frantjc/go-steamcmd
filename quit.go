package steamcmd

var Quit = &q{}

type q struct{}

func (q *q) String() string {
	return "quit"
}

func (*q) Check(_ *Flags) error {
	return nil
}

func (*q) Args() ([]string, error) {
	return []string{"quit"}, nil
}

func (*q) Modify(_ *Flags) error {
	return nil
}
