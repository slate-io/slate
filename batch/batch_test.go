package batch

import (
	"fmt"
	"github.com/rightlag/slate/command"
	"testing"
)

func TestCommandBatch(t *testing.T) {
	ls := command.NewCommand("ls -al")

	b := NewBatchCommand()
	b.Add(ls)
	if b.commands[0] != ls {
		t.Error()
	}
}

func TestCommandsBatch(t *testing.T) {
	pwd := command.NewCommand("pwd")
	ps := command.NewCommand("ps", "-L")
	cs := []*command.Command{
		pwd,
		ps,
	}

	b := NewBatchCommand()
	b.Add(cs)
	if len(b.commands) != 2 {
		t.Error()
	}
	if b.commands[1] != ps {
		t.Error()
	}
}

func TestStringBatch(t *testing.T) {
	ls := command.NewCommand("ls")

	b := NewBatchCommand()
	b.Add("ls")
	if b.commands[0] != ls {
		t.Error()
	}
}

func TestStringsBatch(t *testing.T) {
	cd := command.NewCommand("cd")
	ls := command.NewCommand("ls -al")
	cs := []string{
		"cd",
		"ls -al",
	}

	b := NewBatchCommand()
	b.Add(cs)
	if len(b.commands) != 2 {
		t.Error()
	}
	fmt.Print(b.commands)

	if b.commands[0] != cd {
		t.Error()
	}
	if b.commands[1] != ls {
		t.Error()
	}
}
