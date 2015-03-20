package shell

import (
	"os/exec"
	"strings"
)

// Tool is a base struct for embed
type Tool struct {
	Name string
	Args []string
}

func (c *Tool) String() string {
	return c.Name + " " + strings.Join(c.Args, " ")
}

// Add args
//   Add("-resize", "100x100")
func (c *Tool) Add(strs ...string) *Tool {
	for _, s := range strs {
		c.Args = append(c.Args, s)
	}
	return c
}

// Commands return name and args list
//   Commands() //=> []string{"convert", "-resize", "100x100", "src1.jpg", "output.jpg"}
func (c *Tool) Commands() []string {
	ret := make([]string, 1+len(c.Args))
	ret = append(ret, c.Name)
	ret = append(ret, c.Args...)
	return ret
}

// Command return *exec.Cmd
func (c *Tool) Command() *exec.Cmd {
	return exec.Command(c.Name, c.Args...)
}

// Run command
//   convert.New().Add("src.jpg", "output.jpg").Run()
func (c *Tool) Run() ([]byte, error) {
	return c.Command().CombinedOutput()
}

// Exec is a alias for Run
func (c *Tool) Exec() ([]byte, error) {
	return c.Run()
}

// Spawn command
func (c *Tool) Spawn() {
	go c.Run()
}
