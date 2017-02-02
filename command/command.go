package command

import (
	"fmt"
	"strings"
)

type Command struct {
	args []string
}

func NewCommand(args ...string) *Command {
	return &Command{args}
}

func (c *Command) String() string {
	return strings.Join(c.args, " ")
}

func (c *Command) Execute() error {
	// TODO: Create `execute` functionality.
	fmt.Println(c.String())
	return nil
}
