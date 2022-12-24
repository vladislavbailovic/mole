package internal

import (
	"testing"
	"time"
)

func Test_NewLimitedJob(t *testing.T) {
	dur := time.Second * 4
	job := NewLimitedJob("testdata", []string{}, dur)

	if job.interval != defaultInterval {
		t.Errorf("interval: want %v, got %v",
			defaultInterval,
			job.interval)
	}
}

func Test_NewJob(t *testing.T) {
	job := NewJob("testdata", []string{})

	if job.timeout != nil {
		t.Errorf("duration: want nil, got %v",
			job.timeout)
	}
}
