package internal

import (
	"fmt"
	"os"
	"os/exec"
)

type CommandTarget int

const (
	TargetAny      CommandTarget = iota // Executes always TODO: do we really need this?
	TargetAll                           // Executes if anything changes
	TargetExisting                      // Executes if existing files
	TargetAdded
	TargetChanged
	TargetRemoved
)

type Command struct {
	bin    string
	args   []string
	target CommandTarget
}

func NewCommand(bin string, args []string) *Command {
	return &Command{
		bin:    bin,
		args:   args,
		target: TargetExisting,
	}
}

func (x *Command) SetTarget(tgt CommandTarget) {
	x.target = tgt
}

func (x *Command) ExecuteWith(cmp *Comparison) error {
	var args []string
	if cmp != nil {
		var tgts []string
		switch x.target {
		case TargetExisting:
			tgts = cmp.Existing()
		case TargetAll, TargetAny:
			tgts = cmp.All()
		case TargetAdded:
			tgts = cmp.Added()
		case TargetChanged:
			tgts = cmp.Changed()
		case TargetRemoved:
			tgts = cmp.Removed()
		}
		if len(tgts) == 0 {
			if x.target != TargetAny {
				fmt.Println("no targets, bailing out")
				return nil
			}
		}
		args = make([]string, 0, len(x.args)+len(tgts))
		args = append(args, x.args...)
		args = append(args, tgts...)
	} else {
		args = x.args[:]
	}

	cmd := exec.Command(x.bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}