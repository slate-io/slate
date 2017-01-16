package command

import (
	"testing"
)

func TestNewSshTransport(t *testing.T) {
	c := NewCommand("ls", "-al")
	if c.String() != "ls -al" {
		t.Error()
	}
}
