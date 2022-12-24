package internal

import (
	"testing"
	"time"
)

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
