package main

import (
	"mole/internal"
	"testing"
	"time"
)

func Test_ParseFlags(t *testing.T) {
	es := 11 * time.Second
	suite := map[string]struct {
		args []string
		want Config
	}{
		"ls -la ./testdata/**/*": {
			args: []string{
				"--command", "'ls -la'",
				"--path", "'./testdata/**/*'",
			},
			want: Config{
				Paths:    []string{"./testdata/**/*"},
				Cmd:      "ls -la",
				Interval: internal.DefaultInterval,
			},
		},
		"ls -la ./testdata/**/* every 5s": {
			args: []string{
				"--command", "'ls -la'",
				"--path", "'./testdata/**/*'",
				"--interval", "5s",
			},
			want: Config{
				Paths:    []string{"./testdata/**/*"},
				Cmd:      "ls -la",
				Interval: 5 * time.Second,
			},
		},
		"ls -la ./testdata/**/* every 5s for 11s": {
			args: []string{
				"--command", "'ls -la'",
				"--path", "'./testdata/**/*'",
				"--interval", "5s",
				"--timeout", "11s",
			},
			want: Config{
				Paths:    []string{"./testdata/**/*"},
				Cmd:      "ls -la",
				Interval: 5 * time.Second,
				Timeout:  &es,
			},
		},
	}
	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			compareConfigs(test.want,
				ParseFlags(test.args), t)
		})
	}
}

func Test_JsonConfig(t *testing.T) {
	target := internal.TargetAll
	src := Config{
		Paths:     []string{"./testdata/**/*"},
		Cmd:       "ls -la",
		Interval:  time.Second * 2,
		RawTarget: "all",
		Target:    &target,
	}

	pickle := PickleConfig([]Config{src})
	loaded := UnpickleConfig(pickle)

	compareConfigs(src, loaded[0], t)
}

func compareConfigs(want, got Config, t *testing.T) {
	if got.Paths[0] != want.Paths[0] {
		t.Errorf(
			"paths mismatch: want %+v, got %+v",
			want.Paths, got.Paths)
	}

	if got.Target != nil && want.Target != nil {
		if *got.Target != *want.Target {
			t.Errorf(
				"Target mismatch: want %+v, got %+v",
				*want.Target, *got.Target)
		}
	} else {
		if got.Target != want.Target {
			t.Errorf(
				"Target mismatch: want %+v, got %+v",
				want.Target, got.Target)
		}
	}

	if got.Timeout != nil && want.Timeout != nil {
		if *got.Timeout != *want.Timeout {
			t.Errorf(
				"Timeout mismatch: want %+v, got %+v",
				want.Timeout, got.Timeout)
		}
	} else {
		if got.Timeout != want.Timeout {
			t.Errorf(
				"Timeout mismatch: want %+v, got %+v",
				want.Timeout, got.Timeout)
		}
	}

	if got.Interval != want.Interval {
		t.Errorf(
			"Interval mismatch: want %+v, got %+v",
			want.Interval, got.Interval)
	}
	if got.Maxdepth != want.Maxdepth {
		t.Errorf(
			"Maxdepth mismatch: want %+v, got %+v",
			want.Maxdepth, got.Maxdepth)
	}
	if got.Cmd != want.Cmd {
		t.Errorf(
			"command mismatch: want %+v, got %+v",
			want.Cmd, got.Cmd)
	}
}

func Test_hydrateConfig_Defaults(t *testing.T) {
	c := hydrateConfig(Config{RawErrorHandling: "report"})

	if c.Interval != internal.DefaultInterval {
		t.Errorf("want default interval, got: %v", c.Interval)
	}
	if c.ErrorHandling != internal.OnErrorReport {
		t.Errorf("want report handling, got: %v",
			c.ErrorHandling)
	}
}

func Test_UnpickleRcFile(t *testing.T) {
	cnt := internal.GetTestFile("molerc.json")
	cfgs := UnpickleConfig(cnt)

	if len(cfgs) != 1 {
		t.Errorf("expected one config, got %d", len(cfgs))
	}
}
