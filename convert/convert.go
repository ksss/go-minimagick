package convert

import "github.com/ksss/go-minimagick/shell"

type convert struct {
	shell.Tool
}

// New instance
func New() *convert {
	c := new(convert)
	c.Name = "convert"
	c.Args = []string{}
	return c
}

func (c *convert) Resize(arg string) *convert {
	c.Add("-resize", arg)
	return c
}

func (c *convert) Thumbnail(arg string) *convert {
	c.Add("-thumbnail", arg)
	return c
}

func (c *convert) Define(arg string) *convert {
	c.Add("-define", arg)
	return c
}

func (c *convert) Gravity(arg string) *convert {
	c.Add("-gravity", arg)
	return c
}
