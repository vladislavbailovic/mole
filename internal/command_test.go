package internal

import "testing"

func Test_CommandTarget_New(t *testing.T) {
	cmd := NewCommand("", []string{})
	if cmd.target != TargetExisting {
		t.Errorf("default command target should be existing")
	}
}

func Test_CommandTarget_Set(t *testing.T) {
	suite := map[string]CommandTarget{
		"any":     TargetAny,
		"all":     TargetAll,
		"added":   TargetAdded,
		"changed": TargetChanged,
		"removed": TargetRemoved,
	}
	cmd := NewCommand("", []string{})
	for name, want := range suite {
		t.Run(name, func(t *testing.T) {
			cmd.SetTarget(want)
			if want != cmd.target {
				t.Errorf("want %v, got %v", want, cmd.target)
			}
		})
	}
}
