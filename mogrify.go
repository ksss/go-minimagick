package minimagick

type mogrify struct {
	Tool
}

// New instance
func NewMogrify() *mogrify {
	c := new(mogrify)
	c.Name = "mogrify"
	c.Args = []string{}
	return c
}

func (c *mogrify) Resize(arg string) *mogrify {
	c.Add("-resize", arg)
	return c
}
