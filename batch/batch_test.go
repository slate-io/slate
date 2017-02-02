package batch

import (
	"testing"

	"github.com/rightlag/slate"
	"github.com/rightlag/slate/command"
)

func TestCommandBatch(t *testing.T) {
	b := NewBatchCommand()
	b.Add(command.NewCommand("ls -al", nil))
	if b.Length() != 1 {
		t.Error()
	}
	// If `nil` is the second argument to `NewCommand`,
	// then the command is NOT ignored.
	if b.List()[0].Ignore() {
		t.Error()
	}
}

func TestCommandsBatch(t *testing.T) {
	b := NewBatchCommand()
	b.Add([]*command.Command{
		command.NewCommand("ls -al", slate.Boolean(false)),
		command.NewCommand("rm -rf /", slate.Boolean(true)),
	})
	if b.Length() != 2 {
		t.Error()
	}
	// Explicitly ignore the second command.
	if !b.List()[1].Ignore() {
		t.Error()
	}
}
