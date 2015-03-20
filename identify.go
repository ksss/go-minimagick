package minimagick

type identify struct {
	Tool
}

// New instance
func NewIdentify() *identify {
	c := new(identify)
	c.Name = "identify"
	c.Args = []string{}
	return c
}

func (i *identify) Format(args string) *identify {
	i.Add("-format", args)
	return i
}
