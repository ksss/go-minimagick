package identify

import "github.com/ksss/go-minimagick/shell"

type identify struct {
	shell.Tool
}

// New instance
func New() *identify {
	c := new(identify)
	c.Name = "identify"
	c.Args = []string{}
	return c
}

func (i *identify) Format(args string) *identify {
	i.Add("-format", args)
	return i
}
