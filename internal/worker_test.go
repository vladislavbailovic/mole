package internal

import (
	"testing"
	"time"
)

func Test_ErrorHandlingFromString(t *testing.T) {
	suite := map[string]ErrorHandling{
		"":       OnErrorBreak,
		"wat":    OnErrorBreak,
		"break":  OnErrorBreak,
		"report": OnErrorReport,
		"silent": OnErrorSilent,
	}
	for test, want := range suite {
		t.Run(test, func(t *testing.T) {
			got := ErrorHandlingFromString(test)
			if want != got {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

func Test_ErrorHandling_ToString(t *testing.T) {
	suite := map[string]ErrorHandling{
		"break":  OnErrorBreak,
		"report": OnErrorReport,
		"silent": OnErrorSilent,
	}
	for want, hndl := range suite {
		t.Run(want, func(t *testing.T) {
			got := hndl.String()
			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}

func Test_NewLimitedJob(t *testing.T) {
	dur := time.Second * 4
	job := NewLimitedJob("testdata", &Command{}, dur)

	if job.interval != DefaultInterval {
		t.Errorf("interval: want %v, got %v",
			DefaultInterval,
			job.interval)
	}
}

func Test_NewJob(t *testing.T) {
	job := NewJob("testdata", &Command{})

	if job.timeout != nil {
		t.Errorf("duration: want nil, got %v",
			job.timeout)
	}
}

func Test_SetErrorHandling(t *testing.T) {
	job := NewJob("testdata", &Command{})
	suite := map[string]ErrorHandling{
		"break":  OnErrorBreak,
		"report": OnErrorReport,
		"silent": OnErrorSilent,
	}

	for name, want := range suite {
		t.Run(name, func(t *testing.T) {
			job.SetErrorHandling(want)
			if job.onError != want {
				t.Errorf("want %d, got %d", want, job.onError)
			}
		})
	}
}

func Test_SetInterval(t *testing.T) {
	job := NewJob("testdata", &Command{})
	suite := map[string]struct {
		test time.Duration
		want time.Duration
	}{
		"zero": {
			test: 0,
			want: DefaultInterval,
		},
	}

	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			job.SetInterval(test.test)
			if job.interval != test.want {
				t.Errorf("want %v, got %v",
					test.want, job.interval)
			}
		})
	}
}
