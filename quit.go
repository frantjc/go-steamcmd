package steamcmd

var quit = &q{}

type q struct{}

func (q *q) String() string {
	return "quit"
}

func (*q) check(flags *flags) error {
	return nil
}

func (*q) args() ([]string, error) {
	return []string{"quit"}, nil
}

func (*q) modify(_ *flags) error {
	return nil
}
