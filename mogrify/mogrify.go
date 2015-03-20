package mogrify

import (
	"github.com/ksss/go-minimagick/shell"
)

type mogrify struct {
	shell.Tool
}

// New instance
func New() *mogrify {
	c := new(mogrify)
	c.Name = "mogrify"
	c.Args = []string{}
	return c
}

func (c *mogrify) Resize(arg string) *mogrify {
	c.Add("-resize", arg)
	return c
}
