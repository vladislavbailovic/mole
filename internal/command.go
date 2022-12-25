package internal

import (
	"os"
	"os/exec"
)

type CommandTarget int

const (
	TargetNone CommandTarget = iota // No changed params
	TargetAll                       // Any changes
	TargetExisting
	TargetAdded
	TargetChanged
	TargetRemoved
)

func TargetFromString(s string) CommandTarget {
	switch s {
	case "all":
		return TargetAll
	case "existing":
		return TargetExisting
	case "add", "added":
		return TargetAdded
	case "chg", "changed":
		return TargetChanged
	case "rmv", "removed":
		return TargetRemoved
	}
	return TargetNone
}

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
		case TargetAll, TargetNone:
			tgts = cmp.All()
		case TargetAdded:
			tgts = cmp.Added()
		case TargetChanged:
			tgts = cmp.Changed()
		case TargetRemoved:
			tgts = cmp.Removed()
		}
		if len(tgts) == 0 {
			if x.target != TargetNone {
				// fmt.Println("no targets, bailing out")
				return nil
			}
		}
		args = make([]string, 0, len(x.args)+len(tgts))
		args = append(args, x.args...)
		if x.target != TargetNone {
			args = append(args, tgts...)
		}
	} else {
		if x.target != TargetNone {
			return nil
		}
		args = x.args[:]
	}

	// fmt.Println("Gonna run:", x.bin, args)
	cmd := exec.Command(x.bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
