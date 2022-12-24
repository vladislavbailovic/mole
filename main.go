package main

import (
	"context"
	"fmt"
	"time"
)

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

func NewJob() *Job {
	dur := time.Second * 4
	return &Job{
		command:  []string{"yay"},
		interval: defaultInterval,
		timeout:  &dur,
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	job := NewJob()

	tick := time.NewTicker(24 * time.Second)
	defer cancel()
	go watch(ctx, job)
	<-tick.C
	cancel()

	fmt.Println("all done", job)
}

func watch(ctx context.Context, job *Job) {
	tick := time.NewTicker(job.interval)

	var cancel func()
	if job.timeout != nil {
		ctx, cancel = context.WithTimeout(ctx, *job.timeout)
		defer cancel()
	}

work:
	for {
		fmt.Println(job.command, "on", job.filesets)
		select {
		case <-tick.C:
			fmt.Println("then work some more")
			continue
		case <-ctx.Done():
			fmt.Println("context done")
			// return
			break work
		}
	}

	fmt.Println("work cleanup code")
}
