package command

import (
	"strings"
)

type Command struct {
	args []string
}

func NewCommand(args ...string) *Command {
	return &Command{args}
}

func (c Command) String() string {
	return strings.Join(c.args, " ")
}
