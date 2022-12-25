package internal

import (
	"fmt"
	"testing"
)

func Test_CommandTarget_New(t *testing.T) {
	cmd := NewCommand("", []string{})
	if cmd.target != TargetExisting {
		t.Errorf("default command target should be existing")
	}
}

func Test_CommandTarget_Set(t *testing.T) {
	suite := map[string]CommandTarget{
		"any":     TargetNone,
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

func Test_TargetFromString(t *testing.T) {
	suite := map[string]CommandTarget{
		"wat":      TargetNone,
		"":         TargetNone,
		"none":     TargetNone,
		"all":      TargetAll,
		"existing": TargetExisting,
		"add":      TargetAdded,
		"added":    TargetAdded,
		"chg":      TargetChanged,
		"changed":  TargetChanged,
		"rmv":      TargetRemoved,
		"removed":  TargetRemoved,
	}
	for test, want := range suite {
		t.Run(test, func(t *testing.T) {
			got := TargetFromString(test)
			if want != got {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}

func Test_ExecuteWith(t *testing.T) {
	file := GetTestFilePath("molerc.json")
	suite := map[string]struct {
		target CommandTarget
		test   Filelist
		runs   bool
	}{
		"empty no changes": {
			target: TargetNone,
			test:   NewFilelist([]string{}),
			runs:   true,
		},
		"empty with changes": {
			target: TargetNone,
			test:   NewFilelist([]string{file}),
			runs:   true,
		},
		"add no changes": {
			target: TargetAdded,
			test:   NewFilelist([]string{}),
			runs:   false,
		},
		"add with changes": {
			target: TargetAdded,
			test:   NewFilelist([]string{file}),
			runs:   true,
		},
		"chg no changes": {
			target: TargetChanged,
			test:   NewFilelist([]string{}),
			runs:   false,
		},
		"rmv no changes": {
			target: TargetRemoved,
			test:   NewFilelist([]string{}),
			runs:   false,
		},
		"existing no changes": {
			target: TargetExisting,
			test:   NewFilelist([]string{}),
			runs:   false,
		},
		"existing with changes": {
			target: TargetExisting,
			test:   NewFilelist([]string{file}),
			runs:   true,
		},
	}
	cmd := Command{
		bin:  "nonexistend command nya nya",
		args: []string{"whatever this wont work"},
	}
	old := NewFilelist([]string{})
	for name, test := range suite {
		cmp := CompareFilelists(&test.test, &old)
		t.Run(name, func(t *testing.T) {
			cmd.SetTarget(test.target)
			err := cmd.ExecuteWith(&cmp)
			if test.runs && err == nil {
				t.Error("expected error")
			}
			if !test.runs && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func Test_ExecuteWith_Nil(t *testing.T) {
	cmd := Command{
		bin:  "nonexistend command nya nya",
		args: []string{"whatever this wont work"},
	}
	suite := map[CommandTarget]bool{
		TargetNone:     true,
		TargetAll:      false,
		TargetExisting: false,
		TargetAdded:    false,
		TargetChanged:  false,
		TargetRemoved:  false,
	}
	for tgt, runs := range suite {
		name := fmt.Sprintf("%v runs: %v", tgt, runs)
		t.Run(name, func(t *testing.T) {
			cmd.SetTarget(tgt)
			err := cmd.ExecuteWith(nil)
			if runs && err == nil {
				t.Error("expected error")
			}
			if !runs && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
