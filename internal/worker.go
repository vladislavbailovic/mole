package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var defaultInterval time.Duration = time.Second * 2

type Job struct {
	path     string
	command  []string
	interval time.Duration
	timeout  *time.Duration
	previous *Filelist
}

func NewJob(pathExpr string, command []string) *Job {
	previous := make(Filelist, 0)
	return &Job{
		path:     pathExpr,
		command:  command,
		interval: defaultInterval,
		timeout:  nil,
		previous: &previous,
	}
}

func NewLimitedJob(pathExpr string, commands []string, duration time.Duration) *Job {
	job := NewJob(pathExpr, commands)
	job.timeout = &duration
	return job
}

func (x *Job) SetInterval(i time.Duration) {
	x.interval = i
}

func (x *Job) Watch(ctx context.Context) {
	tick := time.NewTicker(x.interval)

	var cancel func()
	if x.timeout != nil {
		ctx, cancel = context.WithTimeout(ctx, *x.timeout)
		defer cancel()
	}

work:
	for {
		x.execute()
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

func (x *Job) execute() {
	paths := ListFiles(x.path, DefaultGlobDepth)

	lst := NewFilelist(paths)
	cmp := CompareFilelists(&lst, x.previous)
	x.previous = &lst
	if !cmp.Any() {
		fmt.Println("no changes, carry on")
		return
	}

	raw := make([]string, 0, len(x.command)+len(paths))
	raw = append(raw, x.command...)
	raw = append(raw, paths...)
	cmd := exec.Command(raw[0], raw[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
