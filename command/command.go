package command

import (
	"fmt"
	"strings"

	"github.com/rightlag/slate"
)

type Command struct {
	args   []string
	ignore *bool
}

func NewCommand(commandText string, ignore *bool) *Command {
	args := strings.Split(commandText, " ")
	if ignore == nil {
		ignore = slate.Boolean(false)
	}
	return &Command{args, ignore}
}

func (c *Command) Debug() []string {
	return c.args
}

func (c *Command) Ignore() bool {
	return *(c.ignore)
}

func (c *Command) Execute() error {
	// TODO: Create `Execute` functionality.
	fmt.Println(c.Debug())
	return nil
}
