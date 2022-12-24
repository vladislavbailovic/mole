package internal

import "time"

var defaultInterval time.Duration = time.Second * 2

type Fileset struct {
	path string
}

type Job struct {
	filesets []Fileset
	command  []string
	interval time.Duration
	timeout  *time.Duration
}

func NewJob(duration time.Duration) *Job {
	return &Job{
		command:  []string{"yay"},
		interval: defaultInterval,
		timeout:  &duration,
	}
}
