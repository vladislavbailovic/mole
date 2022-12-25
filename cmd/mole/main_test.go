package main

import (
	"mole/internal"
	"testing"
)

func Test_rcfile2JobList(t *testing.T) {
	file := internal.GetTestFilePath("molerc.json")
	jobs := rcfile2JobList(file)

	if len(jobs) != 1 {
		t.Errorf("expected 1 job, got %d", len(jobs))
	}
}
