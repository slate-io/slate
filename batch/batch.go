package batch

import (
	"github.com/rightlag/slate/command"
)

type BatchCommand struct {
	commands []*command.Command
}

func NewBatchCommand() *BatchCommand {
	return &BatchCommand{}
}

func (b *BatchCommand) Add(v interface{}) {
	switch t := v.(type) {
	case *command.Command:
		b.commands = append(b.commands, t)
	case []*command.Command:
		for _, c := range t {
			b.commands = append(b.commands, c)
		}
	case string:
		b.commands = append(b.commands, command.NewCommand(t))
	case []string:
		for _, s := range t {
			b.commands = append(b.commands, command.NewCommand(s))
		}
	default:
		return
	}
}

func (b *BatchCommand) execute() error {
	for _, c := range b.commands {
		if err := c.Execute(); err != nil {
			return err
		}
	}
	return nil
}
