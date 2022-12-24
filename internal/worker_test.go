package internal

import (
	"testing"
	"time"
)

func Test_NewJob(t *testing.T) {
	dur := time.Second * 4
	job := NewJob(dur)

	if job.interval != defaultInterval {
		t.Errorf("interval: want %v, got %v",
			defaultInterval,
			job.interval)
	}
}
