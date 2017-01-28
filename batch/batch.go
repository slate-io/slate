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
	default:
		return
	}
}

func (b *BatchCommand) Length() int {
	return len(b.commands)
}

func (b *BatchCommand) List() []*command.Command {
	return b.commands
}

func (b *BatchCommand) Execute() {
	for _, c := range b.commands {
		if err := c.Execute(); err != nil {
			continue
		}
	}
}
