package command

import (
	"fmt"
	"testing"
)

func TestNewCommand(t *testing.T) {
	c := NewCommand("ls -al", nil)
	fmt.Println(c.Debug())
	if c.Ignore() {
		t.Error()
	}
}
